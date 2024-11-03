// This package contains methods to log data to console and file. It is currently using build-in log/slog package for logging
// In future if we want to replace log/slog with any other package then we only need to modify this file
package log

import "log/slog"

// Global variable to store the slog.Logger object
var logger *slog.Logger

// This function is called only once when its imported and is used to initialize the global variable
func init() {
	logger := newSlogLogger()
}

// Private function that configures log/slog that will log data to both console and file
func newSlogLogger() {
	file, err := os.OpenFile("app.log", os.O_APPEND | os.O_CREATE | os.O_WRONGLY, 0664)
	if err != nil {
		fmt.Printf("Failed to open the log file %v\n", err)
		return
	}
	defer file.Close()

	multiWriter := slog.MultiWriter(os.Stdout, file)

	return slog.New(slog.NewTextHandler(multiWriter))
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