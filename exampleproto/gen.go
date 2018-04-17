package exampleproto

//go:generate protoc  -I=./ -I=$GOPATH/src/weavelab.xyz/insys-onboarding/vendor -I=$GOPATH/src --go_out=plugins=grpc:. example.proto
