package argsmap

import "k8s.io/klog/v2"

type ArgsMap map[string]string

func NewArgsMap() ArgsMap {
	argsmap := make(map[string]string, 0)
	return argsmap
}

func (a ArgsMap) Get(key string) (string, bool) {
	arg, ok := a[key]
	if !ok {
		klog.Infof("key %s not found in argsmap", key)
		return "", false
	}
	return arg, true
}

func (a ArgsMap) Set(key, arg string) {
	a[key] = arg
}
