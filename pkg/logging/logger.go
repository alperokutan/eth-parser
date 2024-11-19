package logging

import (
	"go.uber.org/zap"
)

// Logger wraps zap.SugaredLogger to provide a simple interface for logging.
type Logger struct {
	sugarLogger *zap.SugaredLogger
}

// NewLogger creates and initializes a new Logger instance.
// It uses zap's production configuration for structured logging.
// Logs are written to stdout or a file, depending on the zap configuration.
func NewLogger() *Logger {
	logger, _ := zap.NewProduction() // Use production mode for structured logging
	defer logger.Sync()              // Ensure logs are written to output before exiting
	sugar := logger.Sugar()

	return &Logger{sugarLogger: sugar}
}

// Info logs a message at the Info level.
func (l *Logger) Info(msg string) {
	l.sugarLogger.Info(msg)
}

// Infof logs a formatted message at the Info level.
func (l *Logger) Infof(format string, args ...interface{}) {
	l.sugarLogger.Infof(format, args...)
}

// Error logs a message at the Error level.
func (l *Logger) Error(msg string) {
	l.sugarLogger.Error(msg)
}

// Errorf logs a formatted message at the Error level.
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.sugarLogger.Errorf(format, args...)
}

// Debug logs a message at the Debug level.
// Use this for detailed information during development or troubleshooting.
func (l *Logger) Debug(msg string) {
	l.sugarLogger.Debug(msg)
}

// Debugf logs a formatted message at the Debug level.
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.sugarLogger.Debugf(format, args...)
}

// Warn logs a message at the Warn level.
// Use this for potential issues that are not errors but should be reviewed.
func (l *Logger) Warn(msg string) {
	l.sugarLogger.Warn(msg)
}

// Warnf logs a formatted message at the Warn level.
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.sugarLogger.Warnf(format, args...)
}
