package command

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"ljw/billadm/cmd/options"
	constant "ljw/billadm/const"
	"ljw/billadm/pkg/api/service"
	"ljw/billadm/pkg/controller"
)

var controllers = map[string]controller.Controller{
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
					run(cmd.Parent().Use, cmd.Use, cr.opts)
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
	// 连接grpc服务器
	conn, err := grpc.Dial("unix:///opt/bill/billserver.sock", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 延迟关闭连接
	defer conn.Close()

	// 初始化Greeter服务客户端
	stc := service.NewStorageServiceClient(conn)

	// 初始化上下文，设置请求超时时间为1秒
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// 延迟关闭请求会话
	defer cancel()

	if strings.EqualFold(op, constant.Activate) {
		if _, err := stc.SetCurrentBillName(ctx, &service.SetCurrentBillNameRequest{Name: resource}); err != nil {
			fmt.Printf("activate current bill failed -> <%v>\n", err)
			return
		}
		return
	}

	//验证opts的有效性
	if err := opts.Validate(op, resource); err != nil {
		fmt.Printf("options Validate failed -> <%v>\n", err)
		return
	}

	cfg, err := opts.Config()
	if err != nil {
		fmt.Printf("transfer to Config failed -> <%v>", err)
		return
	}
	fmt.Println(cfg.Time)

	switch op {
	case constant.Get:
		err = controllers[resource].Get(ctx, stc, cfg)
	case constant.Create:
		err = controllers[resource].Create(ctx, stc, cfg)
	case constant.Delete:
		err = controllers[resource].Delete(ctx, stc, cfg)
	case constant.Edit:
		err = controllers[resource].Edit(ctx, stc, cfg)
	default:
		fmt.Printf("op [%s] is invalid\n", op)
	}

	if err != nil {
		fmt.Printf("%s %s failed -> <%v>\n", op, resource, err)
		return
	}
}
