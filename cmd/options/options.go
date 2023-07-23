package options

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"

	constant "ljw/billadm/const"
	v1 "ljw/billadm/pkg/api/v1"
	"ljw/billadm/utils/time"
)

type Options struct {
	// Bill
	name string
	User string

	// day entry 2006-01-02 01-02 02
	// 这个值需要被刷新
	Time string
	All  bool

	// record
	Id          int
	Cost        float32
	Description string
	Label       int
}

func (opt *Options) ApplyTo(fs *pflag.FlagSet) {
	fs.StringVar(&opt.name, "name", "", "specify name for a bill")
	fs.StringVar(&opt.User, "user", "", "specify user for a bill")
	fs.StringVar(&opt.Time, "time", "", "specify time for day entry")
	fs.BoolVar(&opt.All, "all", false, "list all DayEntry")
	fs.IntVar(&opt.Id, "id", -1, "specify id for record")
	fs.Float32Var(&opt.Cost, "cost", -1, "specify cost for record")
	fs.StringVar(&opt.Description, "dsp", "", "specify description for record")
	fs.IntVar(&opt.Label, "label", 0, "specify description for record")
}

func (opt *Options) Validate(op, resource string) error {
	if strings.EqualFold(op, constant.Activate) {
		return nil
	}
	if strings.EqualFold(op, constant.Get) {
		switch resource {
		case constant.Label:
		case constant.Record:
		case constant.DayEntry:
		case constant.Bill:
		default:
			return fmt.Errorf("[%s] not supported", resource)
		}
	} else if strings.EqualFold(op, constant.Create) {
		switch resource {
		case constant.Label:
		case constant.Record:
			if opt.Cost <= 0 {
				return fmt.Errorf("specify cost > 0  for [%s %s]", op, resource)
			}
			if opt.Description == "" {
				return fmt.Errorf("specify description for [%s %s]", op, resource)
			}
			if opt.Id < 0 {
				return fmt.Errorf("specify id for [%s %s]", op, resource)
			}
			if opt.Label > 0 && opt.Label >= len(v1.LabelList) {
				return fmt.Errorf("specify id in [1,%d] for [%s %s]", len(v1.LabelList)-1, op, resource)
			}
		case constant.DayEntry:
		case constant.Bill:
			if opt.name == "" {
				return fmt.Errorf("specify a name for [%s %s]", op, resource)
			}
		default:
			return fmt.Errorf("[%s] not supported", resource)
		}
	} else if strings.EqualFold(op, constant.Delete) {
		switch resource {
		case constant.Label:
		case constant.Record:
			if opt.Id < 0 {
				return fmt.Errorf("specify id for [%s %s]", op, resource)
			}
		case constant.DayEntry:
		case constant.Bill:
			if opt.name == "" {
				return fmt.Errorf("specify a name for [%s %s]", op, resource)
			}
		default:
			return fmt.Errorf("[%s] not supported", resource)
		}
	} else if strings.EqualFold(op, constant.Edit) {
		switch resource {
		case constant.Label:
		case constant.Record:
			if opt.Id < 0 {
				return fmt.Errorf("specify id for [%s %s]", op, resource)
			}
			if opt.Label > 0 && opt.Label >= len(v1.LabelList) {
				return fmt.Errorf("specify id in [1,%d] for [%s %s]", len(v1.LabelList)-1, op, resource)
			}
		case constant.DayEntry:
		case constant.Bill:
			if opt.User == "" {
				return fmt.Errorf("specify user for [%s %s]", op, resource)
			}
		default:
			return fmt.Errorf("[%s] not supported", resource)
		}
	} else {
		return fmt.Errorf("[%s] not supported", op)
	}
	return nil
}

func (opt *Options) Config() (*v1.Config, error) {
	// 刷新时间
	// 可能的输入，09;9;7-9;7-09;07-09;07-9;2013-07-09
	// 合法的输入, 2013-07-09
	// 刷新id
	// 刷新label
	cfg := &v1.Config{
		ID:          fmt.Sprintf("%03d", opt.Id),
		Label:       v1.LabelList[opt.Label],
		Name:        opt.name,
		Cost:        opt.Cost,
		Description: opt.Description,
		User:        opt.User,
		All:         opt.All,
	}

	y, m, d := timeutils.GetYearMonthDay(timeutils.GetNowTimeString())
	t := []string{y, m, d}
	if opt.Time == "" {
		cfg.Time = strings.Join(t, "-")
		return cfg, nil
	}
	tl := strings.Split(opt.Time, "-")
	switch len(tl) {
	case 1:
		t[2] = tl[0]
	case 2:
		t[1] = tl[0]
		t[2] = tl[1]
	case 3:
		t = tl
	default:
		return nil, fmt.Errorf("invalid time input")
	}
	if len(t[1]) == 1 {
		t[1] = "0" + t[1]
	}
	if len(t[2]) == 1 {
		t[2] = "0" + t[2]
	}
	cfg.Time = strings.Join(t, "-")
	return cfg, nil
}
