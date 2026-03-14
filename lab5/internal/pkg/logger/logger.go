package logger

import (
    "fmt"
    "time"
)

// Logger интерфейс для логирования
type Logger interface {
    Info(msg string, args ...interface{})
    Debug(msg string, args ...interface{})
    Error(msg string, args ...interface{})
}

// ConsoleLogger реализует Logger для вывода в консоль
type ConsoleLogger struct {
    debugMode bool
}

// New создает новый консольный логгер
func New(debugMode bool) *ConsoleLogger {
    return &ConsoleLogger{
        debugMode: debugMode,
    }
}

// Info выводит информационное сообщение
func (l *ConsoleLogger) Info(msg string, args ...interface{}) {
    l.print("INFO", msg, args...)
}

// Debug выводит отладочное сообщение (только если debugMode включен)
func (l *ConsoleLogger) Debug(msg string, args ...interface{}) {
    if l.debugMode {
        l.print("DEBUG", msg, args...)
    }
}

// Error выводит сообщение об ошибке
func (l *ConsoleLogger) Error(msg string, args ...interface{}) {
    l.print("ERROR", msg, args...)
}

func (l *ConsoleLogger) print(level, msg string, args ...interface{}) {
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    formattedMsg := fmt.Sprintf(msg, args...)
    fmt.Printf("[%s] %s: %s\n", timestamp, level, formattedMsg)
}
