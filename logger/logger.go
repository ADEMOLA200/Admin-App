package logger

import (
	"log"
	"os"
)

type Logger struct {
	*log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (l *Logger) Info(msg string) {
	l.SetPrefix("INFO: ")
	l.Println(msg)
}

func (l *Logger) Fatal(msg string) {
	l.SetPrefix("FATAL: ")
	l.Println(msg)
	os.Exit(1)
}
