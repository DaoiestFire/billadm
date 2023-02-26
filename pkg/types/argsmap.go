package types

import "fmt"

var _ IArgsMap = &ArgsMap{}

func NewArgsMap() IArgsMap {
	return &ArgsMap{
		args: make(map[string]string, 0),
	}
}

type ArgsMap struct {
	args map[string]string
}

func (am *ArgsMap) Get(key string) (string, error) {
	val, ok := am.args[key]
	if !ok {
		return "", fmt.Errorf("get value from argsmap faild [key:%s]", key)
	}
	return val, nil
}

func (am *ArgsMap) Set(key string, val string) {
	am.args[key] = val
}
