package wgrpcserver

import (
	"context"
	"io"
	"strconv"
	"strings"
	"time"

	"weavelab.xyz/monorail/shared/wlib/wgrpc/wgrpcserver/wrapstream"
	"weavelab.xyz/monorail/shared/wlib/wmetrics"
	"google.golang.org/grpc"
)

const grpcStatsPrefix = "grpc"

type metrics interface {
	Time(time.Duration, string, ...string)
	SetLabels(string, ...string)
}

var (
	WMetricsClient metrics
	grpcLabels     = []string{"endpoint", "direction", "code"}
)

func init() {
	//WMetricsClient that the stats middleware will use to send stats
	WMetricsClient = wmetrics.DefaultClient
	WMetricsClient.SetLabels(grpcStatsPrefix, grpcLabels...)
}

//UnaryStats for stats for unary gRPC endpoints
func UnaryStats(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	resp, err := handler(ctx, req)
	stats(start, info.FullMethod, "unary", err)
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

	//Stats when messages are received and sent
	newStream.RegisterRecvMiddleware(statsStreamMiddleware(statName, "message_recevied"))
	newStream.RegisterSendMiddleware(statsStreamMiddleware(statName, "message_sent"))

	//Initiate the sending/receiving of messages
	err := handler(srv, newStream)

	//Sends stats for stream
	stats(start, info.FullMethod, "stream", err)
	if err != nil {
		return err
	}

	return nil
}

//statsStreamMiddleware returns a middleware func that can be registered on send/recv message of a stream
//it times the length of the message and increments a counter
func statsStreamMiddleware(statName string, direction string) func(inner wrapstream.StreamHandler) wrapstream.StreamHandler {

	return func(inner wrapstream.StreamHandler) wrapstream.StreamHandler {
		mw := func(m interface{}) error {
			streamStart := time.Now()

			err := inner.Stream(m)
			stats(streamStart, statName, direction, err)
			if err != nil {
				return err
			}
			return nil
		}

		return wrapstream.StreamFunc(mw)
	}
}

//stats sends a timer stat and increments a counter; also increments a counter if there was an error
func stats(start time.Time, statName string, direction string, err error) {
	//EOF means the stream is ending and shouldn't be tracked I don't think?
	if err == io.EOF {
		return
	}

	codeStr := strconv.Itoa(int(grpc.Code(err)))

	statName = strings.Replace(statName, ".", "_", -1)

	WMetricsClient.Time(time.Since(start), grpcStatsPrefix, statName, direction, codeStr)
}
