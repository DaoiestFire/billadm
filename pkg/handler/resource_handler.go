package handler

import (
	cmdoptions "ljw/billadm/cmd/options"
)

// ResourceHandler 资源处理器
// 前台的所有请求都会走到资源处理器。根据子命令的调用传入的操作参数（op）来决定需要什么参数和执行什么动作
type ResourceHandler struct {
}

func (rh *RecordHandler) Run(op string, options *cmdoptions.Options) error {
	return nil
}
