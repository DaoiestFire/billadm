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
