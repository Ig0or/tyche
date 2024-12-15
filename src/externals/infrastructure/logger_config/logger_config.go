package logger_config

import (
	"io"
	"log"
	"os"
)

type LoggerConfig struct {
	info  *log.Logger
	err   *log.Logger
	fatal *log.Logger
}

func NewLoggerConfig() *LoggerConfig {
	writer := io.MultiWriter(os.Stdout)

	flags := log.Ldate | log.Ltime | log.Lshortfile

	Logger := &LoggerConfig{
		info:  log.New(writer, "INFO: ", flags),
		err:   log.New(writer, "ERROR: ", flags),
		fatal: log.New(writer, "FATAL: ", flags),
	}

	return Logger
}

func (logger *LoggerConfig) Info(format string, values ...interface{}) {
	logger.info.Printf(format, values...)
}

func (logger *LoggerConfig) Error(format string, values ...interface{}) {
	logger.err.Printf(format, values...)
}

func (logger *LoggerConfig) Fatal(values ...interface{}) {
	logger.fatal.Fatalln(values...)
}
