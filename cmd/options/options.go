package options

import (
	"strings"

	"github.com/spf13/pflag"

	"ljw/billadm/utils/logger"
	"ljw/billadm/utils/time"
)

type Options struct {
	// Bill
	User string

	// day entry 2006-01-02 01-02 02
	// 这个值需要被刷新
	Time string

	// record
	Id          int
	Cost        float32
	Description string
	Label       int
}

func (opt *Options) Refresh() {
	// Options.Time
	currentTime := timeutils.GetNowTimeString()
	currentYear, currentMonth, currentDay := timeutils.GetYearMonthDay(currentTime)

	if len(opt.Time) == 0 {
		opt.Time = strings.Join([]string{currentYear, currentMonth, currentDay}, "-")
		return
	}

	lenOfTime := len(strings.Split(opt.Time, "-"))

	switch lenOfTime {
	case 1:
		opt.Time = strings.Join([]string{currentYear, currentMonth, opt.Time}, "-")
	case 2:
		opt.Time = strings.Join([]string{currentYear, opt.Time}, "-")
	case 3:
	default:
		logger.Errorf("invalid Options.Time [%s]", opt.Time)
	}
}

func (opt *Options) ApplyTo(fs *pflag.FlagSet) {
	fs.StringVar(&opt.User, "user", "", "specify user for a bill")
	fs.StringVar(&opt.Time, "time", "", "specify time for day entry")
	fs.IntVar(&opt.Id, "id", 0, "specify id for record")
	fs.Float32Var(&opt.Cost, "cost", 0, "specify cost for record")
	fs.StringVar(&opt.Description, "dsp", "", "specify description for record")
	fs.IntVar(&opt.Label, "label", 0, "specify description for record")
}

func (opt *Options) Validate(op, resource string) error {

}

func (opt *Options) Config() *Config {

}

type Config string {

}
