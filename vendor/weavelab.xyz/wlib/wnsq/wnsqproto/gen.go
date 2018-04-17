package wnsqproto

//go:generate protoc -I=./ -I=$GOPATH/src --go_out=import_path=wnsqproto:./ tracing.proto
