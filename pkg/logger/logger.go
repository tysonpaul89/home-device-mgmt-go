package logger

import "log/slog"

// Global variable to store the slog.Logger object
var logger *slog.Logger

// This function is called only once when its imported and is used to initialize the global variable
func init() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
}

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
	log *slug.Logger
}

// Public function that returns the Logger object
// When importing call this function to get the object of the Logger
func NewLogger() Logger {
	return loggerImpl{
		log: logger
	}
}

func (l *Logger) Debug(msg string, args ...any) {
	l.log.Debug(msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	l.log.Info(msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.log.Warn(msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.log.Error(msg, args...)
}