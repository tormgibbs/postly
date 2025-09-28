package logger

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
)

type Logger struct {
	slogger *slog.Logger
}

func New(level slog.Level) *Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})
	return &Logger{
		slogger: slog.New(handler),
	}
}

func (l *Logger) Info(msg string, args ...any) {
	l.slogger.Info(msg, args...)
}

func (l *Logger) InfoJSON(msg string, v any) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		l.slogger.Error("failed to marshal JSON", "error", err)
		return
	}
	l.slogger.Info(msg, "payload", string(data))
}

func (l *Logger) Error(msg string, args ...any) {
	l.slogger.Error(msg, args...)
}

func (l *Logger) Fatal(msg string, args ...any) {
	l.slogger.Error(msg, args...)
	os.Exit(1)
}

func (l *Logger) Pretty(msg string, v any) {
	data, _ := json.MarshalIndent(v, "", "  ")
	fmt.Printf("%s:\n%s\n", msg, string(data))
}
