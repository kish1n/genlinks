package handlers

import (
	"context"
	"log"
	"os"
)

// Определение ключа для хранения логгера в контексте
type ctxKeyLog struct{}

// CtxLog возвращает функцию, добавляющую логгер в контекст
func CtxLog(log *log.Logger) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ctxKeyLog{}, log)
	}
}

// LogFromCtx извлекает логгер из контекста
func LogFromCtx(ctx context.Context) *log.Logger {
	if log, ok := ctx.Value(ctxKeyLog{}).(*log.Logger); ok {
		return log
	}
	return log.New(os.Stdout, "default: ", log.LstdFlags)
}
