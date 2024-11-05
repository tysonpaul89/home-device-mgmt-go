// Contains CustomTextHandler and its method definitions
package log

import (
	"context"
	"fmt"
	"io"
	"log"
	"log/slog"
	"time"
)

// CustomTextHandler implements the slog.Handler interface
type CustomTextHandler struct {
	logger *log.Logger
	level  slog.Level
}

// Enabled implements slog.Handler.
func (h *CustomTextHandler) Enabled(context.Context, slog.Level) bool {
	return true
}

// NewCustomTextHandler creates a new CustomTextHandler
func NewCustomTextHandler(out io.Writer, level slog.Level) *CustomTextHandler {
	return &CustomTextHandler{
		logger: log.New(out, "", 0), // No prefix, no default flags
		level:  level,
	}
}

// Handle log entries with custom formatting
func (h *CustomTextHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String()
	msg := r.Message
	now := time.Now()
	formattedTime := fmt.Sprintf("[%s] %s", level, now.Format(time.RFC3339Nano))

	// Log the formatted message
	h.logger.Println(formattedTime, msg)
	return nil
}

// WithAttrs is a no-op for this implementation
func (h *CustomTextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

// WithGroup is a no-op for this implementation
func (h *CustomTextHandler) WithGroup(name string) slog.Handler {
	return h
}
