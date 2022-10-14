package log

import (
	"flag"

	"k8s.io/klog/v2"

	"ljw/billadm/utils/paths"
)

func init() {
	fs := &flag.FlagSet{}
	klog.InitFlags(fs)

	_ = fs.Set("log_file", paths.GetLogFilePath())
	_ = fs.Set("toStderr", "false")
}
