package fileutils

import (
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
