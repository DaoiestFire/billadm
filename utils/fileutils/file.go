package fileutils

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"ljw/billadm/const"
)

func Exist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || os.IsExist(err)
}

func CreateDirectory(path string) error {
	err := os.MkdirAll(path, constant.DirPerm)
	return err
}

func RemoveDirectoryOrFile(filePath string) error {
	return os.RemoveAll(filePath)
}

func IsDirectoryEmpty(path string) bool {
	dirs, _ := os.ReadDir(path)
	if len(dirs) == 0 {
		return true
	}

	return false
}

// RemoveFileRecursive 移除一个文件后，如果目录已经空了就将目录删除
func RemoveFileRecursive(filePath, end string) error {
	if filePath == end {
		return nil
	}
	err := RemoveDirectoryOrFile(filePath)
	if err != nil {
		return err
	}

	parentPath := path.Dir(filePath)
	if IsDirectoryEmpty(parentPath) {
		return RemoveFileRecursive(parentPath, end)
	}
	return nil
}

func ReadFileByte(filePath string) ([]byte, error) {
	if !Exist(filePath) {
		return []byte{}, fmt.Errorf("%s not exist", filePath)
	}

	return os.ReadFile(filePath)
}

func ReadFileString(filePath string) (string, error) {
	buffer, err := ReadFileByte(filePath)
	return string(buffer), err
}

func WriteFileByte(filePath string, data []byte) error {
	fileDir := path.Dir(filePath)
	if !Exist(fileDir) {
		if err := CreateDirectory(fileDir); err != nil {
			return err
		}
	}

	return os.WriteFile(filePath, data, constant.FilePerm)
}

func WriteFileString(filePath string, data string) error {
	return WriteFileByte(filePath, []byte(data))
}

func GenerateJsonData(val any) ([]byte, error) {
	return json.MarshalIndent(val, "", "  ")
}

func FindAllFileFromDirectory(p string) ([]string, error) {
	res := make([]string, 0)
	dirs, err := os.ReadDir(p)
	if err != nil {
		return nil, err
	}
	if len(dirs) == 0 {
		return res, nil
	}
	for _, dir := range dirs {
		nextPath := path.Join(p, dir.Name())
		if dir.IsDir() {
			nextRes, err := FindAllFileFromDirectory(nextPath)
			if err != nil {
				return nil, err
			}
			res = append(res, nextRes...)
		} else {
			res = append(res, nextPath)
		}
	}
	return res, nil
}
