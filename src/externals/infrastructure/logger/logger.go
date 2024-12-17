package logger

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	info  *log.Logger
	err   *log.Logger
	fatal *log.Logger
}

func NewLogger() *Logger {
	writer := io.MultiWriter(os.Stdout)

	flags := log.Ldate | log.Ltime | log.Lshortfile

	Logger := &Logger{
		info:  log.New(writer, "INFO: ", flags),
		err:   log.New(writer, "ERROR: ", flags),
		fatal: log.New(writer, "FATAL: ", flags),
	}

	return Logger
}

func (logger *Logger) Info(format string, values ...interface{}) {
	logger.info.Printf(format, values...)
}

func (logger *Logger) Error(format string, values ...interface{}) {
	logger.err.Printf(format, values...)
}

func (logger *Logger) Fatal(values ...interface{}) {
	logger.fatal.Fatalln(values...)
}
