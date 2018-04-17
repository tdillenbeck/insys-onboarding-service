package wnsq

//go:generate protoc -I=./ -I=$GOPATH/src --go_out=import_path=wnsq:./ tracingprotobuf_test.proto
