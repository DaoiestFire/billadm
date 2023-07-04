package utils

import (
	"strings"

	constant "ljw/billadm/const"
)

func IsResourceValid(resourceName string) bool {
	if strings.EqualFold(constant.Bill, resourceName) ||
		strings.EqualFold(constant.DayEntry, resourceName) ||
		strings.EqualFold(constant.Label, resourceName) ||
		strings.EqualFold(constant.Record, resourceName) {
		return true
	}
	return false
}
