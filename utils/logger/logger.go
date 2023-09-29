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
