package time

import (
	"strings"
	"time"

	constant "ljw/billadm/const"
)

func GetNowTimeString() string {
	return time.Now().Format(constant.TimeFormat)
}

func GetYearFromTimeString(t string) string {
	return strings.Split(t, "-")[0]
}

func GetMonthFromTimeString(t string) string {
	return strings.Split(t, "-")[1]
}

func GetDayFromTimeString(t string) string {
	return strings.Split(t, "-")[2]
}

func GetYearMonthDay(time string) (string, string, string) {
	timeList := strings.Split(time, "-")
	if len(timeList) != 3 {
		return "", "", ""
	}
	return timeList[0], timeList[1], timeList[2]
}
