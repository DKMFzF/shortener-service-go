package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	PATH_LOGGING          = "logs/app.log"
	DEFAULT_LEVEL_LOGIINT = zapcore.ErrorLevel
)

type Logger struct {
	log *zap.SugaredLogger
}

func New(isLogsInFile, isDebugLevel bool) *Logger {
	config := zap.NewProductionConfig()

	if isLogsInFile {
		config.OutputPaths = []string{"stdout", PATH_LOGGING}
	}

	if isDebugLevel {
		config.Level.SetLevel(zapcore.DebugLevel)
	} else {
		/** default logger level */
		config.Level.SetLevel(DEFAULT_LEVEL_LOGIINT)
	}

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
