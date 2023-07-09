package time

import (
	"strings"
	"time"

	constant "ljw/billadm/const"
)

func GetNowTimeString() string {
	return time.Now().Format(constant.TimeFormat)
}

func GetYearMonthDay(time string) (string, string, string) {
	timeList := strings.Split(time, "-")
	if len(timeList) < 3 {
		return "", "", ""
	}
	return timeList[0], timeList[1], timeList[2]
}
