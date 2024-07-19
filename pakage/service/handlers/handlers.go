package handlers

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const logCtxKey ctxKey = iota

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}
