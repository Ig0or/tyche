package i_logger

type ILogger interface {
	Info(format string, values ...interface{})
	Error(format string, values ...interface{})
	Fatal(values ...interface{})
}
