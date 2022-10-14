package types

import "ljw/billadm/utils/argsmap"

var _ Data = &ExecuteData{}

func NewExecuteData(argsmap argsmap.ArgsMap) Data {
	return &ExecuteData{
		ArgsMap: argsmap,
	}
}

type ExecuteData struct {
	ArgsMap argsmap.ArgsMap
}

func (data *ExecuteData) Get(key string) (string, bool) {
	return data.ArgsMap.Get(key)
}
