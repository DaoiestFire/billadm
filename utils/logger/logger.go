package logger

import (
	"flag"

	"k8s.io/klog/v2"
)

func InitLogger(logFile string) {
	flagSet := flag.NewFlagSet("klog", flag.ContinueOnError)
	klog.InitFlags(flagSet)
	flagSet.Set("log_file", logFile)
	flagSet.Set("logtostderr", "false")
	flagSet.Set("stderrthreshold", "3")
}

func Flush() {
	klog.Flush()
}

func Info(args ...interface{}) {
	klog.Info(args...)
}

func Infof(format string, args ...interface{}) {
	klog.Infof(format, args...)
}

func Error(args ...interface{}) {
	klog.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	klog.Errorf(format, args...)
}
