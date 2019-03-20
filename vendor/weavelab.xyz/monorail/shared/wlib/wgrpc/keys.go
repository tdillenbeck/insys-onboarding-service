package wgrpc

import (
	"context"
)

const (
	DebugIDMetadataKey   = "weavedebugid"
	RequestIDMetadataKey = "request_id"
	TokenMetadataKey     = "token"
)

func UserAgent(ctx context.Context) string {
	ua, ok := IncomingMetadata(ctx, "user-agent")
	if !ok {
		return "not_set"
	}
	return ua
}
