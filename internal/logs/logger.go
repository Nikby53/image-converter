package logs

import (
	"fmt"
	"runtime"

	log "github.com/sirupsen/logrus"
)

type StandardLogger struct {
	logger *log.Logger
}

func NewLogger() *StandardLogger {
	var baseLogger = log.New()
	var standardLogger = &StandardLogger{baseLogger}
	standardLogger.logger.SetReportCaller(true)
	standardLogger.logger.Formatter = &log.TextFormatter{
		TimestampFormat:        "2006-01-02 15:04:05",
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			return "", fmt.Sprintf("%s:%d", frame.File, frame.Line)
		},
	}
	return standardLogger
}

func (l *StandardLogger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *StandardLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *StandardLogger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *StandardLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}
