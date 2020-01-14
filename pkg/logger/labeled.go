package logger

type Labeled struct {
	label  string
	logger Logger
}

var _ Logger = &Labeled{}

func (l *Labeled) WithLabel(label string) Logger {
	return &Labeled{label, l}
}

func (l *Labeled) Debug(msg string, args ...interface{}) {
	l.logger.Debug(l.label+": "+msg, args...)
}

func (l *Labeled) Info(msg string, args ...interface{}) {
	l.logger.Info(l.label+": "+msg, args...)
}

func (l *Labeled) Error(msg string, args ...interface{}) {
	l.logger.Error(l.label+": "+msg, args...)
}

func (l *Labeled) Flush() {
	l.logger.Flush()
}
