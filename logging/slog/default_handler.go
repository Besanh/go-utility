package slog

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"path/filepath"
	"runtime"
)

type defaultHandler struct {
	slog.Handler
}

func NewDefaultHandler(w io.Writer, formatter Formatter, opts *slog.HandlerOptions) *defaultHandler {
	if opts == nil {
		opts = &slog.HandlerOptions{}
	}
	var handler slog.Handler
	switch formatter {
	case FORMAT_TEXT:
		handler = slog.NewTextHandler(w, opts)
	default:
		handler = slog.NewJSONHandler(w, opts)
	}
	return &defaultHandler{
		handler,
	}
}

func (h defaultHandler) Handle(ctx context.Context, r slog.Record) error {
	if logId, ok := ctx.Value(keyTraceId).(string); ok {
		r.Add("trace_id", logId)
	} else {
		r.Add("trace_id", "unknown")
	}
	_, path, numLine, _ := runtime.Caller(6)
	srcFile := filepath.Base(path)
	r.Add(slog.String("file", fmt.Sprintf("%s:%d", srcFile, numLine)))
	return h.Handler.Handle(ctx, r)
}
