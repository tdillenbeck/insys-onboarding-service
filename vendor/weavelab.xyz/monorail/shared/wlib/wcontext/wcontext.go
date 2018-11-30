package wcontext

import (
	"context"

	"weavelab.xyz/monorail/shared/wlib/uuid"
)

type keys int

const requestIDKey = keys(0)

//NewWithRequestID adds the id to the context or genereates a new ID if it is empty
func NewWithRequestID(ctx context.Context, id string) context.Context {
	if id == "" {
		id = uuid.NewV4().String()
	}

	return context.WithValue(ctx, requestIDKey, id)
}

//RequestID retrieves the RequestID from the context
func RequestID(ctx context.Context) string {
	r, ok := ctx.Value(requestIDKey).(string)
	if ok == false {
		return ""
	}

	return r
}
