package logger

import (
	"fmt"
	"time"
)

const (
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
)

const (
	LogFormat  = "%s [%s] {%s}"
	TimeFormat = "2006-01-02 15:04:05"
)

func Info(args ...interface{}) {
	msg := fmt.Sprint(args...)
	fmt.Printf(LogFormat, INFO, time.Now().Format(TimeFormat), msg)
}

func Infof(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Printf(LogFormat, INFO, time.Now().Format(TimeFormat), msg)
}

func Warn(args ...interface{}) {
	msg := fmt.Sprint(args...)
	fmt.Printf(LogFormat, WARN, time.Now().Format(TimeFormat), msg)
}

func Warnf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Printf(LogFormat, WARN, time.Now().Format(TimeFormat), msg)
}

func Error(args ...interface{}) {
	msg := fmt.Sprint(args...)
	fmt.Printf(LogFormat, ERROR, time.Now().Format(TimeFormat), msg)
}

func Errorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Printf(LogFormat, ERROR, time.Now().Format(TimeFormat), msg)
}
