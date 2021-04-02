package logger

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type Root struct {
	out io.WriteCloser
	mx  sync.RWMutex
}

var _ Logger = &Root{}

func NewRoot(out io.WriteCloser) *Root {
	return &Root{out: out}
}

func NewFileRoot(path string) (*Root, error) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return nil, err
	}

	return NewRoot(f), nil
}

func MustNewFileRootTruncated(path string) *Root {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		panic(err.Error())
	}

	return NewRoot(f)
}

func MustNewFileRoot(path string) *Root {
	logger, err := NewFileRoot(path)
	if err != nil {
		panic(err.Error())
	}
	return logger
}

func (l *Root) WithLabel(label string) Logger {
	return &Labeled{label, l}
}

func (l *Root) Debug(msg string, args ...interface{}) {
	l.mx.Lock()
	defer l.mx.Unlock()
	l.out.Write([]byte(fmt.Sprintf("debug: "+msg+"\n", args...)))
}

func (l *Root) Info(msg string, args ...interface{}) {
	l.mx.Lock()
	defer l.mx.Unlock()
	l.out.Write([]byte(fmt.Sprintf("info: "+msg+"\n", args...)))
}

func (l *Root) Error(msg string, args ...interface{}) {
	l.mx.Lock()
	defer l.mx.Unlock()
	l.out.Write([]byte(fmt.Sprintf("error: "+msg+"\n", args...)))
}

func (l *Root) Flush() {
	l.mx.Lock()
	defer l.mx.Unlock()
	l.out.Close()
}
