package logger

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	info *log.Logger
	err  *log.Logger
}

func NewLogger() *Logger {
	writer := io.MultiWriter(os.Stdout)

	flags := log.Ldate | log.Ltime

	infoPrefix := "\033[32m[INFO]:\033[0m "
	errorPrefix := "\033[31m[ERROR]:\033[0m "

	Logger := &Logger{
		info: log.New(writer, infoPrefix, flags),
		err:  log.New(writer, errorPrefix, flags),
	}

	return Logger
}

func (logger *Logger) Info(format string, values ...interface{}) {
	logger.info.Printf(format, values...)
}

func (logger *Logger) Error(message string, err error) {
	var errorInString string

	if err != nil {
		errorInString = err.Error()
	}

	logger.err.Printf("%s - Original Error: %s", message, errorInString)
}
