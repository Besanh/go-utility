package slog

import (
	"bytes"
	"strings"
	"testing"
)

func TestDefaultLogger(t *testing.T) {
	buf := new(bytes.Buffer)

	logger := NewDefaultLogger()

	tests := []struct {
		name     string
		args     string
		want     string
		isOrigin bool
	}{
		{
			name:     "Test origin log",
			args:     "log from origin slog",
			want:     "log from origin slog",
			isOrigin: true,
		},
		{
			name:     "Test trace log",
			args:     "log from origin slog",
			want:     "log from origin slog",
			isOrigin: false,
		},
	}

	logger.SetOutput(buf)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isOrigin {
				logger.Info(tt.args)
			} else {
				Info(tt.args)
			}
			if !strings.Contains(buf.String(), tt.want) {
				t.Errorf("TestDefaultLogger() = %v, want %v", buf.String(), tt.want)
			}
		})
	}
}
