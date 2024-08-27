package slog

import (
	"context"
	"fmt"
	"log/slog"
)

// get format msg
func getMessage(template string, fmtArgs []interface{}) string {
	if len(fmtArgs) == 0 {
		return template
	}

	if template != "" {
		return fmt.Sprintf(template, fmtArgs...)
	}

	if len(fmtArgs) == 1 {
		if str, ok := fmtArgs[0].(string); ok {
			return str
		}
	}
	return fmt.Sprint(fmtArgs...)
}

// Adapt level to slog level
func tranSLevel(level Level) (lvl slog.Level) {
	switch level {
	case LEVEL_DEBUG:
		lvl = slog.LevelDebug
	case LEVEL_INFO:
		lvl = slog.LevelInfo
	case LEVEL_WARN:
		lvl = slog.LevelWarn
	case LEVEL_ERROR:
		lvl = slog.LevelError
	default:
		lvl = slog.LevelDebug
	}
	return
}

type ContextKey string

const keyTraceId = ContextKey("trace_id")

func WithTraceId(ctx context.Context, traceId string) context.Context {
	return context.WithValue(ctx, keyTraceId, traceId)
}
