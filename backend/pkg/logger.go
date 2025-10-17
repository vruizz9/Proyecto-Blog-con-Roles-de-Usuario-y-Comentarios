package pkg

import (
	"log"
	"os"
	"time"
)

// Logger proporciona funcionalidad de logging para la aplicación
type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	warnLogger  *log.Logger
}

// NewLogger crea una nueva instancia del logger
func NewLogger() *Logger {
	flags := log.LstdFlags | log.Lshortfile

	return &Logger{
		infoLogger:  log.New(os.Stdout, "[INFO] ", flags),
		errorLogger: log.New(os.Stderr, "[ERROR] ", flags),
		warnLogger:  log.New(os.Stdout, "[WARN] ", flags),
	}
}

// Info registra un mensaje informativo
func (l *Logger) Info(format string, v ...interface{}) {
	l.infoLogger.Printf(format, v...)
}

// Error registra un mensaje de error
func (l *Logger) Error(format string, v ...interface{}) {
	l.errorLogger.Printf(format, v...)
}

// Warn registra un mensaje de advertencia
func (l *Logger) Warn(format string, v ...interface{}) {
	l.warnLogger.Printf(format, v...)
}

// Fatal registra un mensaje fatal y termina la aplicación
func (l *Logger) Fatal(format string, v ...interface{}) {
	l.errorLogger.Printf(format, v...)
	os.Exit(1)
}

// WithTimestamp agrega timestamp al mensaje
func (l *Logger) WithTimestamp(format string, v ...interface{}) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	return timestamp + " " + format
}
