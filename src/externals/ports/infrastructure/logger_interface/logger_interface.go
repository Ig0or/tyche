package logger_interface

type LoggerInterface interface {
	Info(format string, values ...interface{})
	Error(message string, err error)
}
