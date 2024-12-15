package i_logger_config

type ILoggerConfig interface {
	Info(format string, values ...interface{})
	Error(format string, values ...interface{})
	Fatal(values ...interface{})
}
