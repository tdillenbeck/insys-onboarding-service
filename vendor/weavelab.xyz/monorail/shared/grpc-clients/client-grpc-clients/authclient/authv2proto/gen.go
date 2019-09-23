package authv2proto

//go:generate protoc -I=./ -I=$GOPATH/src --go_out=plugins=grpc:. auth.proto
