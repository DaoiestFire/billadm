package fileutils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"ljw/billadm/const"
)

func Exist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || os.IsExist(err)
}

func RemoveFile(filePath string) error {
	return os.RemoveAll(filePath)
}

func ReadFileByte(filePath string) ([]byte, error) {
	if !Exist(filePath) {
		return []byte{}, fmt.Errorf("%s is not exist", filePath)
	}

	return os.ReadFile(filePath)
}

func ReadFileString(filePath string) (string, error) {
	buffer, err := ReadFileByte(filePath)
	return string(buffer), err
}

func WriteFileByte(filePath string, data []byte) error {
	fileDir := filepath.Dir(filePath)
	if !Exist(filepath.Dir(filePath)) {
		err := os.MkdirAll(fileDir, constant.DirPerm)
		if err != nil {
			return fmt.Errorf("mkdir %s failed", fileDir)
		}
	}

	return os.WriteFile(filePath, data, constant.FilePerm)
}

func WriteFileString(filePath string, data string) error {
	return WriteFileByte(filePath, []byte(data))
}

func GenerateJsonData(val any) ([]byte, error) {
	return json.MarshalIndent(val, "  ", "  ")
}
