package models

import "github.com/sirupsen/logrus"

//Структура для логгирования
type Logger struct {
	stdOut *logrus.Logger
	stdErr *logrus.Logger
}

func NewLogger(stdOut *logrus.Logger, stdErr *logrus.Logger) *Logger {

	return &Logger{
		stdOut: stdOut,
		stdErr: stdErr,
	}
}

//nolint
func (l *Logger) Info(args ...interface{}) {
	l.stdOut.Info(args)
}

//nolint
func (l *Logger) Error(args ...interface{}) {
	l.stdErr.Error(args)
}
