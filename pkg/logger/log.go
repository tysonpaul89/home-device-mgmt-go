package logger

import "log/slog"

// Interface containing functions that can be used to log data.
// To use logger you have to call NewLogger() function of this package.
// Example:
// import "hrdm/logger"
// log := logger.NewLogger()
// log.Info("This is an info log")
// NOTE: We are followin this convention because in future if we want to replace log/slog with another logger
// then we only need to modify this file
type Logger interface {
	Debug() nil
	Info() nil
	Warn() nil
	Error() nil
} 


// Private sturct that implements the Logger class
type loggerImpl struct {
}

// Public function that returns the Logger object
// When importing call this function to get the object of the Logger
func NewLogger() Logger {
	return loggerImpl{}
}

func (l *Logger) Debug(msg string, args ...any) {
	slog.Debug(msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	slog.Info(msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	slog.Warn(msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	slog.Error(msg, args...)
}