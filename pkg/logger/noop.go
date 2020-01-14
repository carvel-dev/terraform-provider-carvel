package logger

type Noop struct{}

var _ Logger = Noop{}

func NewNoop() Noop { return Noop{} }

func (l Noop) WithLabel(label string) Logger         { return NewNoop() }
func (l Noop) Debug(msg string, args ...interface{}) {}
func (l Noop) Info(msg string, args ...interface{})  {}
func (l Noop) Error(msg string, args ...interface{}) {}
func (l Noop) Flush()                                {}
