package paths

import (
	"os"
	"path/filepath"

	constant "ljw/billadm/const"
)

func GetSoftwarePath() string {
	return filepath.Dir(os.Args[0])
}

func GetDataPath() string {
	return filepath.Join(GetSoftwarePath(), constant.DataDir)
}

func GetLogFilePath() string {
	return filepath.Join(GetSoftwarePath(), constant.LogDir, constant.LogFile)
}
