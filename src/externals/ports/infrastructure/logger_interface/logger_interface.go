package logger_interface

type LoggerInterface interface {
	Info(format string, values ...interface{})
	Error(format string, values ...interface{})
	Fatal(values ...interface{})
}
