package web

import (
	"context"

	"github.com/google/uuid"
)

type ctxKey int

const (
	traceIDKey ctxKey = iota + 1
)

func setTraceID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, traceIDKey, id)
}

func GetTraceID(ctx context.Context) string {
	v, ok := ctx.Value(traceIDKey).(string)
	if !ok {
		return uuid.UUID{}.String()
	}

	return v
}
