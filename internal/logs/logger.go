package logs

import (
	"runtime"

	log "github.com/sirupsen/logrus"
)

// Logger is struct that holds custom logger.
type Logger struct {
	logger *log.Logger
}

// NewLogger is constructor for Logger.
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

// Infoln logs a message at level Info.
func (l *Logger) Infoln(args ...interface{}) {
	l.logger.Infoln(args...)
}

// Fatalf logs a message at level Fatal.
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

// Infof logs a message at level Info.
func (l *Logger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

// Errorf logs a message at level Error.
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

// Printf logs a message at level Info.
func (l *Logger) Printf(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}
