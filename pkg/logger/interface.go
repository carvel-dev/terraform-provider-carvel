package logger

type Logger interface {
	WithLabel(label string) Logger
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Flush()
}
