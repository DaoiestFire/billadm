package options

import (
	"ljw/billadm/utils/logger"
	timeutils "ljw/billadm/utils/time"
	"strings"
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
}

func (op *Options) Refresh() {
	// Options.Time
	currentTime := timeutils.GetNowTimeString()
	currentYear := timeutils.GetYearFromTimeString(currentTime)
	currentMonth := timeutils.GetMonthFromTimeString(currentTime)
	currentDay := timeutils.GetDayFromTimeString(currentTime)

	lenOfTime := len(strings.Split(op.Time, "-"))

	switch lenOfTime {
	case 0:
		op.Time = strings.Join([]string{currentYear, currentMonth, currentDay}, "-")
	case 1:
		op.Time = strings.Join([]string{currentYear, currentMonth, op.Time}, "-")
	case 2:
		op.Time = strings.Join([]string{currentYear, op.Time}, "-")
	case 3:
	default:
		logger.Errorf("invalid Options.Time [%s]", op.Time)
	}
}
