package logger

import (
	"flag"

	"k8s.io/klog"
)

func InitLogger(logFile string) {
	flagSet := flag.NewFlagSet("klog", flag.ExitOnError)
	klog.InitFlags(flagSet)
	_ = flagSet.Set("log_file", logFile)
}

func Flush() {
	klog.Flush()
}

func Info(args ...interface{}) {
	klog.Info(args)
}

func Infof(format string, args ...interface{}) {
	klog.Infof(format, args)
}

func Warn(args ...interface{}) {
	klog.Warning(args)
}

func Warnf(format string, args ...interface{}) {
	klog.Warningf(format, Warn)
}

func Error(args ...interface{}) {
	klog.Error(args)
}

func Errorf(format string, args ...interface{}) {
	klog.Errorf(format, args)
}
