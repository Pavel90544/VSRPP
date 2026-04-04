package logger

import (
    "fmt"
    "time"
)

type logger struct{}

func New() *logger {
    return &logger{}
}

func (l *logger) Info(msg string) {
    fmt.Printf("[INFO] %s, message - %s\n", time.Now().Format(time.RFC3339), msg)
}

func (l *logger) Debug(msg string) {
    fmt.Printf("[DEBUG] %s, message - %s\n", time.Now().Format(time.RFC3339), msg)
}

func (l *logger) Error(msg string, err error) {
    errMsg := ""
    if err != nil {
        errMsg = " err - " + err.Error()
    }
    fmt.Printf("[ERROR] %s, message - %s%s\n", time.Now().Format(time.RFC3339), msg, errMsg)
}
