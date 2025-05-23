package sl

import (
	"context"
	"log/slog"
)

// NewDiscardLogger returns a logger that discards all messages.
func NewDiscardLogger() *slog.Logger {
	return slog.New(NewDiscardHandler())
}

// DiscardHandler is a slog.Handler that discards all messages.
type DiscardHandler struct{}

// NewDiscardHandler returns a new DiscardHandler.
func NewDiscardHandler() *DiscardHandler {
	return &DiscardHandler{}
}

// Handle discards all messages.
func (h *DiscardHandler) Handle(_ context.Context, _ slog.Record) error {
	return nil
}

// WithAttrs discards all messages.
func (h *DiscardHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return h
}

// WithGroup discards all messages.
func (h *DiscardHandler) WithGroup(_ string) slog.Handler {
	return h
}

// Enabled returns false.
func (h *DiscardHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return false
}
