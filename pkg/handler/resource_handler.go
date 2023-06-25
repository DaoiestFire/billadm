package handler

import (
	cmdoptions "ljw/billadm/cmd/options"
	constant "ljw/billadm/const"
	"ljw/billadm/pkg/manager"
	"ljw/billadm/utils/logger"
)

func NewResourceHandler() *ResourceHandler {
	return &ResourceHandler{}
}

// ResourceHandler 资源处理器
// 前台的所有请求都会走到资源处理器。根据子命令的调用传入的操作参数（op）来决定需要什么参数和执行什么动作
type ResourceHandler struct {
	cm *manager.ConfigManager

	billHandler     IResource
	recordHandler   IResource
	labelHandler    IResource
	dayEntryHandler IResource

	resources Resources
}

// Resources 将所有的资源处理函数再保存一遍。传给每一个资源的处理器
// 这样每一种资源可以方便的使用其他资源的处理器
type Resources struct {
	billHandler     *BillHandler
	recordHandler   *RecordHandler
	labelHandler    *LabelHandler
	dayEntryHandler *DayEntryHandler
}

// IResource op,资源库，配置控制器，前台选项
type IResource interface {
	Run(string, Resources, *manager.ConfigManager, *cmdoptions.Options) error
}

// Run 每一种资源对外的函数都是Run函数。resource Handler接收到资源处理请求中将请求分配给每一种资源处理器
func (rh *ResourceHandler) Run(op, resourceName string, options *cmdoptions.Options) {
	logger.Infof("===start to [%s] resource [%s]===", op, resourceName)

	var err error
	switch resourceName {
	case constant.Bill:
		err = rh.billHandler.Run(op, rh.resources, rh.cm, options)
	case constant.Record:
		err = rh.recordHandler.Run(op, rh.resources, rh.cm, options)
	case constant.DayEntry:
		err = rh.dayEntryHandler.Run(op, rh.resources, rh.cm, options)
	case constant.Label:
		err = rh.labelHandler.Run(op, rh.resources, rh.cm, options)
	default:
		logger.Errorf("invalid resource name [%s]", resourceName)
	}

	if err != nil {
		logger.Infof("[%s] resource [%s] failed", op, resourceName)
	}

	logger.Infof("===end to [%s] resource [%s]===", op, resourceName)
}
