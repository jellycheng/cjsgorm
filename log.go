package cjsgorm

import (
	"fmt"
)

type WriterLogger interface {
	Printf(string, ...interface{})
}

type DefaultLogger struct{}

func NewDefaultLogger() *DefaultLogger {
	return &DefaultLogger{}
}

func (l DefaultLogger) Printf(format string, values ...interface{}) {
	val := fmt.Sprintf(format, values...)
	fmt.Println(val)
}

var MyGormLogObj WriterLogger = DefaultLogger{}
