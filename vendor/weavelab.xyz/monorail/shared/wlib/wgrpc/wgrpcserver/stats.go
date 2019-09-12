package wgrpcserver

import (
	"context"
	"io"
	"net"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/stats"
	"weavelab.xyz/monorail/shared/wlib/wgrpc"
	"weavelab.xyz/monorail/shared/wlib/wgrpc/wgrpcserver/wrapstream"
	"weavelab.xyz/monorail/shared/wlib/wmetrics"
)

const (
	grpcStatsPrefix     = "grpc"
	grpcStatsConnPrefix = "grpc_conn"
)

type metrics interface {
	Time(time.Duration, string, ...string)
	Incr(int, string, ...string)
	SetLabels(string, ...string)
}

var (
	WMetricsClient metrics
	grpcLabels     = []string{"endpoint", "direction", "code", "userAgent", "localAddr"}
)

func init() {
	//WMetricsClient that the stats middleware will use to send stats
	WMetricsClient = wmetrics.DefaultClient
	WMetricsClient.SetLabels(grpcStatsPrefix, grpcLabels...)
	WMetricsClient.SetLabels(grpcStatsConnPrefix, "connAction", "remoteAddr", "localAddr")
}

//UnaryStats for stats for unary gRPC endpoints
func UnaryStats(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	resp, err := handler(ctx, req)
	sendStats(ctx, start, info.FullMethod, "unary", err)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

//StreamStats for handling stats for streaming gRPC endpoints
func StreamStats(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	start := time.Now()
	statName := info.FullMethod

	//Wrap the stream so we can add stats for every message sent and received
	newStream := wrapstream.WrapServerStream(ss)

	ctx := ss.Context()

	//Stats when messages are received and sent
	newStream.RegisterRecvMiddleware(statsStreamMiddleware(ctx, statName, "message_received"))
	newStream.RegisterSendMiddleware(statsStreamMiddleware(ctx, statName, "message_sent"))

	//Initiate the sending/receiving of messages
	err := handler(srv, newStream)

	//Sends stats for stream
	sendStats(ctx, start, info.FullMethod, "stream", err)
	if err != nil {
		return err
	}

	return nil
}

//statsStreamMiddleware returns a middleware func that can be registered on send/recv message of a stream
//it times the length of the message and increments a counter
func statsStreamMiddleware(ctx context.Context, statName string, direction string) func(inner wrapstream.StreamHandler) wrapstream.StreamHandler {

	return func(inner wrapstream.StreamHandler) wrapstream.StreamHandler {
		mw := func(m interface{}) error {
			streamStart := time.Now()

			err := inner.Stream(m)
			sendStats(ctx, streamStart, statName, direction, err)
			if err != nil {
				return err
			}
			return nil
		}

		return wrapstream.StreamFunc(mw)
	}
}

//stats sends a timer stat and increments a counter; also increments a counter if there was an error
func sendStats(ctx context.Context, start time.Time, statName string, direction string, err error) {
	//EOF means the stream is ending and shouldn't be tracked I don't think?
	if err == io.EOF {
		return
	}

	codeStr := strconv.Itoa(int(grpc.Code(err)))

	statName = strings.Replace(statName, ".", "_", -1)

	userAgent := strings.Replace(wgrpc.UserAgent(ctx), ".", "_", -1)

	localAddr := localAddr(ctx)

	WMetricsClient.Time(time.Since(start), grpcStatsPrefix, statName, direction, codeStr, userAgent, localAddr)
}

type statsHandler struct {
}

type contextKey string

var (
	localAddrKey  contextKey = "localAdd"
	remoteAddrKey contextKey = "remoteAddr"
)

// TagRPC can attach some information to the given context.
// The context used for the rest lifetime of the RPC will be derived from
// the returned context.
func (s *statsHandler) TagRPC(ctx context.Context, tag *stats.RPCTagInfo) context.Context {
	return ctx
}

// HandleRPC processes the RPC stats.
func (s *statsHandler) HandleRPC(ctx context.Context, stats stats.RPCStats) {}

// TagConn can attach some information to the given context.
// The returned context will be used for stats handling.
// For conn stats handling, the context used in HandleConn for this
// connection will be derived from the context returned.
// For RPC stats handling,
//  - On server side, the context used in HandleRPC for all RPCs on this
// connection will be derived from the context returned.
//  - On client side, the context is not derived from the context returned.
func (s *statsHandler) TagConn(ctx context.Context, tag *stats.ConnTagInfo) context.Context {
	ctx = context.WithValue(ctx, localAddrKey, tag.LocalAddr)
	ctx = context.WithValue(ctx, remoteAddrKey, tag.RemoteAddr)
	return ctx
}

const unknownAddr = "unknown"

// HandleConn processes the Conn stats.
func (s *statsHandler) HandleConn(ctx context.Context, connStats stats.ConnStats) {
	remote := unknownAddr
	remoteAddr, ok := ctx.Value(remoteAddrKey).(net.Addr)
	if ok {
		remote = remoteAddr.String()
	}

	local := localAddr(ctx)

	// Label values may contain any Unicode characters.
	// wmetrics treats '.' as a separator
	remote = strings.Replace(remote, ".", "_", -1)

	switch connStats.(type) {
	case *stats.ConnBegin:
		WMetricsClient.Incr(1, grpcStatsConnPrefix, "begin", remote, local)
	case *stats.ConnEnd:
		WMetricsClient.Incr(1, grpcStatsConnPrefix, "end", remote, local)
	}
}

// localAddr extracts the local server address from the context
func localAddr(ctx context.Context) string {
	local := unknownAddr
	localAddr, ok := ctx.Value(localAddrKey).(net.Addr)
	if ok {
		local = localAddr.String()
	}
	return local
}
