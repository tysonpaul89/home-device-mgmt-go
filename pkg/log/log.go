// This package contains methods to log data to console and file. It is currently using build-in log/slog package for logging
// In future if we want to replace log/slog with any other package then we only need to modify this file
package log

import (
	"io"
	"log/slog"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

// Global variable to store the slog.Logger object
var logger *slog.Logger

// This function is called only once when its imported and is used to initialize the global variable
func init() {
	logger = newSlogLogger()
}

// Private function that configures log/slog that will log data to both console and file
func newSlogLogger() *slog.Logger {
	// Using lumberjack package to implement log rotation to avoid memory issues
	// Lumberjack handles file closure automatically.
	logFile := &lumberjack.Logger{
		Filename:   "app.log",
		MaxSize:    10,   // Max size in MB before rotation
		MaxBackups: 5,    // Max number of old log files to keep
		MaxAge:     28,   // Max days to retain old log files
		Compress:   true, // Compress rotated log files
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	return slog.New(slog.NewTextHandler(multiWriter, &slog.HandlerOptions{}))
}

func Debug(msg string, args ...any) {
	logger.Debug(msg, args...)
}

func Info(msg string, args ...any) {
	logger.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	logger.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	logger.Error(msg, args...)
}
