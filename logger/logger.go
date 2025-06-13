package logger

import (
	logrus "github.com/sirupsen/logrus"
)

type Logger struct {
	logger *logrus.Logger
}

func NewLogger() Logger {
	var logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	return Logger{logger: logger}
}

func (l *Logger) Request(request map[string]interface{}) {
	l.logger.WithFields(request).Info("Request")
}

func (l *Logger) Response(response map[string]interface{}) {
	l.logger.WithFields(response).Info("Response")
}

func (l *Logger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *Logger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}
