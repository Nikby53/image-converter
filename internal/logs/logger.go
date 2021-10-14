package logs

import (
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
			return "", ""
		},
	}
	return standardLogger
}

func (l *StandardLogger) Infoln(args ...interface{}) {
	l.logger.Infoln(args...)
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

func (l *StandardLogger) Printf(format string, args ...interface{}) {
	l.logger.Printf(format, args)
}
