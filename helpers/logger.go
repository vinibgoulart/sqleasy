package helpers

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	*log.Logger
}

func (l *Logger) Info(message string) {
	l.Printf("INFO: %s", message)
}

func (l *Logger) Error(message string) {
	l.Printf("ERROR: %s", message)
}

func (l *Logger) Debug(message string) {
	l.Printf("DEBUG: %s", message)
}

func LoggerCreate(context string) *Logger {
	logger := log.New(os.Stdout, fmt.Sprintf("%s: ", context), log.LstdFlags)
	return &Logger{logger}
}
