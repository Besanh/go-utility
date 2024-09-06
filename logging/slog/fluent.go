package slog

import (
	"context"
	"fmt"
	"log/slog"
)

type FluentLogger struct {
	l      *slog.Logger
	config *config
}

func NewFluentLogger(opts ...Option) *FluentLogger {
	config := defaultConfig()
	for _, opt := range opts {
		opt.apply(config)
	}
	// When user set the handlerOptions level but not set with coreconfig level
	if !config.coreConfig.withLevel && config.coreConfig.withHandlerOptions && config.coreConfig.opt.Level != nil {
		lvl := &slog.LevelVar{}
		lvl.Set(config.coreConfig.opt.Level.Level())
		config.coreConfig.level = lvl
	}
	config.coreConfig.opt.Level = config.coreConfig.level
	config.fluentConfig.level = config.coreConfig.level
	logger := slog.New(NewFluentdHandler(config.coreConfig.writer, config.coreConfig.opt, config.fluentConfig))
	return &FluentLogger{
		l:      logger,
		config: config,
	}
}

var _ ILogger = (*FluentLogger)(nil)

func (l *FluentLogger) Log(level Level, msg string) {
	logger := l.l.With()
	logger.Log(context.TODO(), tranSLevel(level), msg)
}

func (l *FluentLogger) Logf(level Level, format string, kvs ...interface{}) {
	logger := l.l.With()
	msg := getMessage(format, kvs)
	logger.Log(context.TODO(), tranSLevel(level), msg)
}

func (l *FluentLogger) LogCtxf(level Level, ctx context.Context, format string, kvs ...interface{}) {
	logger := l.l.With()
	msg := getMessage(format, kvs)
	logger.Log(ctx, tranSLevel(level), msg)
}

func (l *FluentLogger) LogWithArgs(level Level, ctx context.Context, msg string, args ...interface{}) {
	logger := l.l.With(args...)
	logger.Log(ctx, tranSLevel(level), msg)
}

func (l *FluentLogger) Debug(args ...any) {
	l.Log(LEVEL_DEBUG, fmt.Sprint(args...))
}

func (l *FluentLogger) Info(args ...any) {
	l.Log(LEVEL_INFO, fmt.Sprint(args...))
}

func (l *FluentLogger) Warn(args ...any) {
	l.Log(LEVEL_WARN, fmt.Sprint(args...))
}

func (l *FluentLogger) Error(args ...any) {
	l.Log(LEVEL_ERROR, fmt.Sprint(args...))
}

func (l *FluentLogger) Debugf(msg string, args ...any) {
	l.Logf(LEVEL_DEBUG, msg, args...)
}

func (l *FluentLogger) Infof(msg string, args ...any) {
	l.Logf(LEVEL_INFO, msg, args...)
}

func (l *FluentLogger) Warnf(msg string, args ...any) {
	l.Logf(LEVEL_WARN, msg, args...)
}

func (l *FluentLogger) Errorf(msg string, args ...any) {
	l.Logf(LEVEL_ERROR, msg, args...)
}

func (l *FluentLogger) DebugContext(ctx context.Context, args ...any) {
	l.LogCtxf(LEVEL_DEBUG, ctx, fmt.Sprint(args...))
}

func (l *FluentLogger) InfoContext(ctx context.Context, args ...any) {
	l.LogCtxf(LEVEL_INFO, ctx, fmt.Sprint(args...))
}

func (l *FluentLogger) WarnContext(ctx context.Context, args ...any) {
	l.LogCtxf(LEVEL_WARN, ctx, fmt.Sprint(args...))
}

func (l *FluentLogger) ErrorContext(ctx context.Context, args ...any) {
	l.LogCtxf(LEVEL_ERROR, ctx, fmt.Sprint(args...))
}

func (l *FluentLogger) DebugfContext(ctx context.Context, msg string, args ...any) {
	l.LogCtxf(LEVEL_DEBUG, ctx, msg, args...)
}

func (l *FluentLogger) InfofContext(ctx context.Context, msg string, args ...any) {
	l.LogCtxf(LEVEL_INFO, ctx, msg, args...)
}

func (l *FluentLogger) WarnfContext(ctx context.Context, msg string, args ...any) {
	l.LogCtxf(LEVEL_WARN, ctx, msg, args...)
}

func (l *FluentLogger) ErrorfContext(ctx context.Context, msg string, args ...any) {
	l.LogCtxf(LEVEL_ERROR, ctx, msg, args...)
}

func (l *FluentLogger) SetLevel(level Level) {
	lvl := tranSLevel(level)
	l.config.coreConfig.level.Set(lvl)
}
