package logger

import (
	"fmt"
	"log"

	"github.com/CrazyThursdayV50/gotils/pkg/logger"
)

type loggerImpl struct{}

func log_(level string, format string, a ...any) {
	log.Output(3, fmt.Sprintf("[%s] %s", level, fmt.Sprintf(format, a...)))
}

func (l *loggerImpl) Debug(f string, a ...any) {
	log_("DEBUG", f, a...)
}

func (l *loggerImpl) Info(f string, a ...any) {
	log_("INFO", f, a...)
}

func (l *loggerImpl) Warn(f string, a ...any) {
	log_("WARN", f, a...)
}

func (l *loggerImpl) Error(f string, a ...any) {
	log_("ERROR", f, a...)
}

func DefaultLogger() logger.Logger {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile | log.Lmsgprefix)
	return &loggerImpl{}
}
