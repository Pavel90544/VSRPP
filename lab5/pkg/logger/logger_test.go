package logger

import (
	"errors"
	"testing"
)

func TestLogger_Info(t *testing.T) {
	l := New()
	l.Info("Test info message")
}

func TestLogger_Debug(t *testing.T) {
	l := New()
	l.Debug("Test debug message")
}

func TestLogger_Error(t *testing.T) {
	l := New()
	err := errors.New("something went wrong")
	l.Error("Test error", err) // теперь передаём настоящую ошибку
}

func TestLogger_ErrorWithNil(t *testing.T) {
	l := New()
	l.Error("Test error with nil", nil) // nil не вызовет паники
}

func TestLogger_New(t *testing.T) {
	l := New()
	if l == nil {
		t.Error("New() returned nil")
	}
}
