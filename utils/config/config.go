package configutils

import (
	"os"
	"path/filepath"
)

var InstallPath string

func init() {
	currentBinPath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	realBinPath, err := filepath.EvalSymlinks(currentBinPath)
	if err != nil {
		panic(err)
	}
	realBinDir := filepath.Dir(realBinPath)
	InstallPath = filepath.Dir(realBinDir)
}
