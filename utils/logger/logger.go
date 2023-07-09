package logger

import (
	"k8s.io/klog"
)

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
