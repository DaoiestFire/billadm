package options

import (
	"strings"

	"github.com/spf13/pflag"

	"ljw/billadm/utils/logger"
	timeutils "ljw/billadm/utils/time"
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

func (op *Options) Refresh() {
	// Options.Time
	currentTime := timeutils.GetNowTimeString()
	currentYear, currentMonth, currentDay := timeutils.GetYearMonthDay(currentTime)

	if len(op.Time) == 0 {
		op.Time = strings.Join([]string{currentYear, currentMonth, currentDay}, "-")
		return
	}

	lenOfTime := len(strings.Split(op.Time, "-"))

	switch lenOfTime {
	case 1:
		op.Time = strings.Join([]string{currentYear, currentMonth, op.Time}, "-")
	case 2:
		op.Time = strings.Join([]string{currentYear, op.Time}, "-")
	case 3:
	default:
		logger.Errorf("invalid Options.Time [%s]", op.Time)
	}
}

func (op *Options) ApplyTo(fs *pflag.FlagSet) {
	fs.StringVar(&op.User, "user", "", "specify user for a bill")
	fs.StringVar(&op.Time, "time", "", "specify time for day entry")
	fs.IntVar(&op.Id, "id", 0, "specify id for record")
	fs.Float32Var(&op.Cost, "cost", 0, "specify cost for record")
	fs.StringVar(&op.Description, "dsp", "", "specify description for record")
	fs.IntVar(&op.Label, "label", 0, "specify description for record")
}
