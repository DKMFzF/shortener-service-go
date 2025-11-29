package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	log *zap.SugaredLogger
}

func (l *Logger) New() *Logger {
	config := zap.NewProductionConfig()

	config.OutputPaths = []string{"stdout", "logs/app.log"}
	config.Level.SetLevel(zapcore.ErrorLevel)
	config.Encoding = "json"

	rawLogger, err := config.Build()
	if err != nil {
		return nil
	}

	return &Logger{
		log: rawLogger.Sugar(),
	}
}

func (l *Logger) Infof(forStr string, args ...any) {
	l.log.Infof(forStr, args...)
}

func (l *Logger) Debugf(forStr string, args ...any) {
	l.log.Debugf(forStr, args...)
}

func (l *Logger) Fatalf(forStr string, args ...any) {
	l.log.Fatalf(forStr, args...)
}

func (l *Logger) Warnf(forStr string, args ...any) {
	l.log.Warnf(forStr, args...)
}
