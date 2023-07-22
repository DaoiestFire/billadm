package command

import (
	"fmt"
	v1 "ljw/billadm/pkg/api/v1"
	"ljw/billadm/pkg/controller"
	"strings"

	"github.com/spf13/cobra"

	"ljw/billadm/cmd/options"
	constant "ljw/billadm/const"
	"ljw/billadm/pkg/storage"
	"ljw/billadm/utils/logger"
)

var controllers = map[string]v1.Controller{
	constant.Bill:     &controller.BillController{},
	constant.Label:    &controller.LabelController{},
	constant.DayEntry: &controller.DayEntryController{},
	constant.Record:   &controller.RecordController{},
}

// 基于工厂模式实现子命令的注册

type SubCommandRegister struct {
	opSubCommand       []string
	resourceSubCommand []string
	opts               *options.Options
}

func NewCommandRegister() *SubCommandRegister {
	cr := &SubCommandRegister{
		resourceSubCommand: []string{constant.Record, constant.DayEntry, constant.Label, constant.Bill},
		opSubCommand:       []string{constant.Get, constant.Create, constant.Delete, constant.Edit},
		opts:               &options.Options{},
	}
	return cr
}

func (cr *SubCommandRegister) BindToCommand(command *cobra.Command) {
	for _, op := range cr.opSubCommand {
		opCommand := &cobra.Command{
			Use: op,
		}
		for _, resource := range cr.resourceSubCommand {
			resourceCommand := &cobra.Command{
				Use: resource,
				Run: func(cmd *cobra.Command, args []string) {
					run(op, resource, cr.opts)
				},
			}
			cr.opts.ApplyTo(resourceCommand.Flags())
			opCommand.AddCommand(resourceCommand)
		}
		command.AddCommand(opCommand)
	}
	command.AddCommand(NewActivateCommand())
}

func NewActivateCommand() *cobra.Command {
	command := &cobra.Command{
		Use: constant.Activate,
		Run: func(cmd *cobra.Command, args []string) {
			var billName string
			if len(args) == 0 {
				billName = ""
			} else {
				billName = args[0]
			}
			run(constant.Activate, billName, nil)
		},
		Args: cobra.RangeArgs(0, 1),
	}
	return command
}

func run(op, resource string, opts *options.Options) {
	// 获取storage
	st, err := storage.GetStorage()
	if err != nil {
		logger.Errorf("GetStorage failed -> <%v>", err)
		fmt.Printf("GetStorage failed -> <%v>\n", err)
		return
	}
	defer func() {
		err := st.Finalizer()
		if err != nil {
			logger.Errorf("storage finalizer failed -> <%v>", err)
			fmt.Printf("storage finalizer failed -> <%v>", err)
		}
	}()

	//验证opts的有效性
	err = opts.Validate(op, resource)
	if err != nil {
		logger.Errorf("options Validate failed -> <%v>", err)
		fmt.Printf("options Validate failed -> <%v>\n", err)
		return
	}

	cfg := opts.Config()

	if strings.EqualFold(op, constant.Activate) {
		err = st.SetCurrentBillName(resource)
		if err != nil {
			logger.Errorf("activate current bill failed -> <%v>", err)
			fmt.Printf("activate current bill failed -> <%v>\n", err)
			return
		}
		return
	}

	switch op {
	case constant.Get:
		err = controllers[resource].Get(st, cfg)
	case constant.Create:
		err = controllers[resource].Create(st, cfg)
	case constant.Delete:
		err = controllers[resource].Delete(st, cfg)
	case constant.Edit:
		err = controllers[resource].Edit(st, cfg)
	default:
		logger.Errorf("op [%s] is invalid", op)
		fmt.Printf("op [%s] is invalid\n", op)
	}

	if err != nil {
		logger.Errorf("%s %s failed -> <%v>", op, resource, err)
		fmt.Printf("%s %s failed -> <%v>", op, resource, err)
		return
	}
}
