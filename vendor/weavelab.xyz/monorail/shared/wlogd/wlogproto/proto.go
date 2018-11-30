package wlogproto

//go:generate protoc -I=. --go_out=plugins=grpc:$GOPATH/src wlog.proto
