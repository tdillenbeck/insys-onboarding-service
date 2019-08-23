package wlog

import (
	"context"
	"fmt"
	"net"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	"weavelab.xyz/monorail/shared/wlib/wlog/tag"
	"weavelab.xyz/monorail/shared/wlogd/chunker"
	"weavelab.xyz/monorail/shared/wlogd/wlogproto"
)

var port string

func init() {
	port = os.Getenv("WLOG_PORT")

	if port == "" {
		port = "3050"
	}

	l, _ := strconv.ParseBool(os.Getenv("WLOG_LOCAL"))
	AlwaysLogLocally(l)

	wlogdChan = make(chan wlogQueueMsg, 30)
	go wlogdConnect()
}

type wlogQueueMsg struct {
	c     context.Context
	mtype LogMsgType
	msg   string
	tags  []tag.Tag
	file  string
	line  int
}

var (
	alwaysLogLocally = false

	wlogdChan      chan wlogQueueMsg
	wlogdConnected bool // only set by wlogdConnectOne
)

func WlogdLogger(c context.Context, mtype LogMsgType, msg string, tags []tag.Tag) {

	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "???"
		line = 0
	}

	select {
	case wlogdChan <- wlogQueueMsg{c, mtype, msg, tags, file, line}:
	default:
	}

	if alwaysLogLocally || !wlogdConnected {
		stdoutLogger(c, mtype, msg, tags, file, line)
	}

}

func wlogdConnect() {
	waitTime := 1
	for {
		connected, _ := wlogdConnectOne()

		// simple backoff
		if connected {
			waitTime = 1
		} else if waitTime < 5 {
			waitTime += 1
		}

		time.Sleep(time.Duration(waitTime) * time.Second)
	}
}

func wlogdConnectOne() (bool, error) {
	addr, err := net.ResolveTCPAddr("tcp", net.JoinHostPort("localhost", port))
	if err != nil {
		return false, err
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		return false, err
	}
	defer func() {
		wlogdConnected = false
		_ = conn.Close()
	}()

	wlogdConnected = true

	_ = conn.SetKeepAlivePeriod(1 * time.Minute)
	_ = conn.SetKeepAlive(true)

	// Use a buffer instead of proto.Marshal to better reuse memory
	buf := proto.NewBuffer(make([]byte, 0, 4096))

	cw := chunker.NewChunkWriter(conn)

	// write the magic WHOAMI message
	buf.Reset()
	err = buf.Marshal(&whoAmILog)
	if err != nil {
		return true, err
	}

	err = cw.WriteChunk(buf.Bytes())
	if err != nil {
		return true, err
	}

	for {
		queueMsg := <-wlogdChan
		log := formatMsg(queueMsg)

		buf.Reset()
		err := buf.Marshal(&log)

		if err != nil {
			StdoutLogger(queueMsg.c, queueMsg.mtype, queueMsg.msg, queueMsg.tags)
			log2 := wlogproto.Log{Message: err.Error()}
			buf.Reset()
			err2 := buf.Marshal(&log2)
			if err2 != nil {
				fmt.Println("Error marshaling!!! ", err, err2)
				continue
			}
		}

		err = cw.WriteChunk(buf.Bytes())
		if err != nil {
			file := queueMsg.file
			line := int(queueMsg.line)

			stdoutLogger(queueMsg.c, queueMsg.mtype, queueMsg.msg, queueMsg.tags, file, line)
			return true, err
		}
	}
}

func formatMsg(queueMsg wlogQueueMsg) wlogproto.Log {
	_, mtype, msg, tags := queueMsg.c, queueMsg.mtype, queueMsg.msg, queueMsg.tags

	var log wlogproto.Log

	log.Timestamp = time.Now().Unix()

	switch mtype {
	case ERROR:
		log.Level = wlogproto.Level_ERROR
	case WARN:
		log.Level = wlogproto.Level_WARN
	case INFO:
		log.Level = wlogproto.Level_INFO
	case DEBUG:
		log.Level = wlogproto.Level_DEBUG
	case TRACE:
		log.Level = wlogproto.Level_TRACE
	}

	log.Message = msg

	// WError messages
	if mtype == ERROR && msg == "" && len(tags) >= 1 && tags[0].Type == tag.WErrorType && tags[0].Key == "" {
		werr := tags[0].WErrorVal

		log.Message = werr.Message()
		extraMsgs := werr.ExtraMessages()
		if len(extraMsgs) > 0 {
			log.Message += "\n" + strings.Join(extraMsgs, "\n")
		}
	}

	log.TagsString = make(map[string]string)
	log.TagsInt = make(map[string]int32)
	log.TagsInt64 = make(map[string]int64)
	log.TagsFloat = make(map[string]float32)
	log.TagsBool = make(map[string]bool)
	log.TagsDuration = make(map[string]*wlogproto.Duration)

	for _, t := range tags {
		switch t.Type {
		case tag.StringType:
			log.TagsString[t.Key] = t.StringVal
		case tag.IntType:
			log.TagsInt64[t.Key] = t.IntVal
		case tag.DurationType:
			log.TagsDuration[t.Key] = &wlogproto.Duration{Duration: int64(t.DVal)}
		case tag.FloatType:
			log.TagsFloat[t.Key] = float32(t.FloatVal)
		case tag.BoolType:
			log.TagsBool[t.Key] = t.BoolVal
		case tag.WErrorType:
			// Message should be handled above for wlog.WError messages
			// just extract tags and stacktrace here
			werr := t.WErrorVal
			if t.Key != "" {
				log.TagsString[t.Key] = werr.Message()
			}

			stack := werr.Stack()

			st := make([]*wlogproto.StackEntry, len(stack))

			for i, e := range stack {
				st[i] = &wlogproto.StackEntry{
					File: e.File,
					Name: e.Name,
					Line: int32(e.Line),
				}
			}

			log.StackTrace = st

			tags := werr.Tags()
			for key, val := range tags {
				switch v := val.(type) {
				case string:
					log.TagsString[key] = v
				case int:
					log.TagsInt[key] = int32(v)
				case int32:
					log.TagsInt[key] = v
				case int64:
					log.TagsInt64[key] = v
				case float32:
					log.TagsFloat[key] = v
				case float64:
					log.TagsFloat[key] = float32(v)
				case bool:
					log.TagsBool[key] = v
					// case error: <= Not sure exactly how to handle this...
				default:
					log.TagsString[key] = fmt.Sprint(v)
				}
			}
		}
	}

	log.TagsString["_file"] = queueMsg.file
	log.TagsInt["_line"] = int32(queueMsg.line)

	// if no stacktrace was provided, add one of our own
	if len(log.StackTrace) == 0 {
		// TODO: add stacktrace
	}

	return log
}

func AlwaysLogLocally(l bool) {
	alwaysLogLocally = l
}
