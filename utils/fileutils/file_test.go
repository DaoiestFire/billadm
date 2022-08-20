package fileutils

import (
	"fmt"
	"testing"
)

func TestExsit(t *testing.T) {
	if !Exist("./file.go") {
		t.Fail()
	}
}

func TestNoExist(t *testing.T) {
	if Exist("./test.go") {
		t.Fail()
	}
}

func TestReadFileString(t *testing.T) {
	res, err := ReadFileString("../../testing/test.log")
	if err != nil {
		t.Fail()
	}
	fmt.Println(res)
}
