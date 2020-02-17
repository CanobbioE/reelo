package infrastructure

import (
	"io"
	"log"
	"os"
)

func init() {
	// Configure Logging
	// logFilePath := "log.txt"
	logFilePath := os.Getenv("LOG_FILE_PATH")
	if logFilePath != "" {
		f, err := os.OpenFile("./"+logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
		wrt := io.MultiWriter(os.Stdout, f)
		log.SetOutput(wrt)
	}
}

// Logger represents a logger
type Logger struct {
}

// NewLogger is the Loger contructor
func NewLogger() *Logger {
	return &Logger{}
}

// Log logs a message to the configured output stream
func (l *Logger) Log(msg string, a ...interface{}) {
	log.Printf(msg, a)
}
