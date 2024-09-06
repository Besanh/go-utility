package slog

import (
	"context"
	"io"

	"log/slog"

	"github.com/fluent/fluent-logger-golang/fluent"
	slogcommon "github.com/samber/slog-common"
)

type FluentConfig struct {
	// log level (default: debug)
	level slog.Leveler

	// connection to Fluentd
	Client *fluent.Fluent
	Tag    string

	// configal: customize json payload builder
	Converter Converter
	// configal: fetch attributes from context
	AttrFromContext []func(ctx context.Context) []slog.Attr

	// configal: see slog.HandlerOptions
	AddSource   bool
	ReplaceAttr func(groups []string, a slog.Attr) slog.Attr
}

type FluentdHandler struct {
	attrs  []slog.Attr
	groups []string
	config FluentConfig
	slog.Handler
}

func NewFluentdHandler(w io.Writer, opts *slog.HandlerOptions, cfg FluentConfig) slog.Handler {
	if cfg.Client == nil {
		panic("missing Fuentd client")
	}

	if cfg.Converter == nil {
		cfg.Converter = DefaultConverter
	}

	if cfg.AttrFromContext == nil {
		cfg.AttrFromContext = []func(ctx context.Context) []slog.Attr{}
	}

	return &FluentdHandler{
		Handler: slog.NewJSONHandler(w, opts),
		config:  cfg,
	}
}

var _ slog.Handler = (*FluentdHandler)(nil)

func (h *FluentdHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.config.level.Level()
}

func (h *FluentdHandler) Handle(ctx context.Context, record slog.Record) error {
	if err := h.postToFluent(ctx, record); err != nil {
		return err
	}
	return h.Handler.Handle(ctx, record)
}

func (h *FluentdHandler) postToFluent(ctx context.Context, record slog.Record) error {
	tag := h.getTag(&record)
	fromContext := ContextExtractor(ctx, h.config.AttrFromContext)
	message := h.config.Converter(h.config.AddSource, h.config.ReplaceAttr, append(h.attrs, fromContext...), h.groups, &record, tag)

	return h.config.Client.PostWithTime(tag, record.Time, message)
}

func (h *FluentdHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &FluentdHandler{
		config: h.config,
		attrs:  slogcommon.AppendAttrsToGroup(h.groups, h.attrs, attrs...),
		groups: h.groups,
	}
}

func (h *FluentdHandler) WithGroup(name string) slog.Handler {
	// https://cs.opensource.google/go/x/exp/+/46b07846:slog/handler.go;l=247
	if name == "" {
		return h
	}

	return &FluentdHandler{
		config: h.config,
		attrs:  h.attrs,
		groups: append(h.groups, name),
	}
}

func (h *FluentdHandler) getTag(record *slog.Record) string {
	tag := h.config.Tag

	for i := range h.attrs {
		if h.attrs[i].Key == "tag" && h.attrs[i].Value.Kind() == slog.KindString {
			tag = h.attrs[i].Value.String()
			break
		}
	}

	record.Attrs(func(attr slog.Attr) bool {
		if attr.Key == "tag" && attr.Value.Kind() == slog.KindString {
			tag = attr.Value.String()
			return false
		}
		return true
	})

	return tag
}
