package slog

import (
	"context"
	"fmt"
)

var logger IFullLogger = NewDefaultLogger()

func Info(args ...any) {
	logger.
		Info(fmt.Sprint(args...))
}

func Warn(args ...any) {
	logger.
		Warn(fmt.Sprint(args...))
}

func Error(args ...any) {
	logger.
		Error(fmt.Sprint(args...))
}

func Debug(args ...any) {
	logger.
		Debug(fmt.Sprint(args...))
}

func Infof(msg string, args ...any) {
	logger.
		Info(fmt.Sprintf(msg, args...))
}

func Warnf(msg string, args ...any) {
	logger.
		Warn(fmt.Sprintf(msg, args...))
}

func Errorf(msg string, args ...any) {
	logger.
		Error(fmt.Sprintf(msg, args...))
}

func Debugf(msg string, args ...any) {
	logger.
		Debug(fmt.Sprintf(msg, args...))
}

func InfoContext(ctx context.Context, args ...any) {
	logger.
		InfoContext(ctx, fmt.Sprint(args...))
}

func WarnContext(ctx context.Context, args ...any) {
	logger.
		WarnContext(ctx, fmt.Sprint(args...))
}

func ErrorContext(ctx context.Context, args ...any) {
	logger.
		ErrorContext(ctx, fmt.Sprint(args...))
}

func DebugContext(ctx context.Context, args ...any) {
	logger.
		DebugContext(ctx, fmt.Sprint(args...))
}

func InfofContext(ctx context.Context, msg string, args ...any) {
	logger.
		InfoContext(ctx, fmt.Sprintf(msg, args...))
}

func WarnfContext(ctx context.Context, msg string, args ...any) {
	logger.
		WarnContext(ctx, fmt.Sprintf(msg, args...))
}

func ErrorfContext(ctx context.Context, msg string, args ...any) {
	logger.
		ErrorContext(ctx, fmt.Sprintf(msg, args...))
}

func DebugfContext(ctx context.Context, msg string, args ...any) {
	logger.
		DebugContext(ctx, fmt.Sprintf(msg, args...))
}
