package web

import (
	"fmt"
	"time"
)

// LogCategory represents different categories of log messages
type LogCategory string

const (
	LogInit    LogCategory = "INIT"
	LogConfig  LogCategory = "CONFIG"
	LogSEO     LogCategory = "SEO"
	LogI18N    LogCategory = "I18N"
	LogRouter  LogCategory = "ROUTER"
	LogRender  LogCategory = "RENDER"
	LogIndex   LogCategory = "INDEX"
	LogCheck   LogCategory = "CHECK"
	LogServer  LogCategory = "SERVER"
	LogMonitor LogCategory = "MONITOR"
	LogError   LogCategory = "ERROR"
	LogWarn    LogCategory = "WARN"
)

// Logger provides standardized logging with timestamps and categories
type Logger struct {
	startTime time.Time
}

// NewLogger creates a new logger instance
func NewLogger() *Logger {
	return &Logger{startTime: time.Now()}
}

// Log outputs a standardized log message with optional duration
func (l *Logger) Log(category LogCategory, emoji, action, description string, duration ...time.Duration) {
	timestamp := time.Now().Format("15:04:05.000")
	categoryStr := fmt.Sprintf("%-8s", string(category))
	
	if len(duration) > 0 {
		ms := float64(duration[0].Nanoseconds()) / 1e6
		fmt.Printf("[%s] [%s] %s %s: %s (took %.3fms)\n", 
			timestamp, categoryStr, emoji, action, description, ms)
	} else {
		fmt.Printf("[%s] [%s] %s %s: %s\n", 
			timestamp, categoryStr, emoji, action, description)
	}
}

// GetStartTime returns the logger's start time
func (l *Logger) GetStartTime() time.Time {
	return l.startTime
}