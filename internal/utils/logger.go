package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

// Logger represents a basic logger with different log levels.
type Logger struct {
	LogLevel LogLevel
}

// LogLevel represents the severity level of a log message.
type LogLevel int

const (
	Info LogLevel = iota
	Warning
	Error
)

var logColorFuncs = map[LogLevel]*color.Color{
	Info:    color.New(color.FgGreen),
	Warning: color.New(color.FgYellow),
	Error:   color.New(color.FgRed),
}

// NewLogger creates a new Logger with the specified log level.
func NewLogger(logLevel LogLevel) *Logger {
	return &Logger{
		LogLevel: logLevel,
	}
}

// log formats and prints a log message with the appropriate color.
func (l *Logger) log(level LogLevel, format string, args ...interface{}) {
	if l.LogLevel > level {
		return // Log level not enabled
	}
	colorFunc, exists := logColorFuncs[level]
	if !exists {
		colorFunc = color.New()
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	message := fmt.Sprintf(format, args...)
	formattedMessage := fmt.Sprintf("[%s] %s: %s\n", timestamp, strings.ToUpper(fmt.Sprintf("%d", level)), message)
	colorFunc.Printf(formattedMessage)
}

// Info logs an informational message.
func (l *Logger) Info(format string, args ...interface{}) {
	l.log(Info, format, args...)
}

// Warning logs a warning message.
func (l *Logger) Warning(format string, args ...interface{}) {
	l.log(Warning, format, args...)
}

// Error logs an error message.
func (l *Logger) Error(format string, args ...interface{}) {
	l.log(Error, format, args...)
}

// String returns the string representation of a LogLevel.
func (l LogLevel) String() string {
	switch l {
	case Info:
		return "info"
	case Warning:
		return "warning"
	case Error:
		return "error"
	default:
		return "unknown"
	}
}
