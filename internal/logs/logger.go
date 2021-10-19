package logs

import (
	"runtime"

	log "github.com/sirupsen/logrus"
)

type Logger struct {
	logger *log.Logger
}

func NewLogger() *Logger {
	var baseLogger = log.New()
	var standardLogger = &Logger{baseLogger}
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

func (l *Logger) Infoln(args ...interface{}) {
	l.logger.Infoln(args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *Logger) Printf(format string, args ...interface{}) {
	l.logger.Printf(format, args)
}
