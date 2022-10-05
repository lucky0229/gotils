package string

import (
	"strconv"
)

func IsEmpty(s string) bool {
	return len(s) == 0
}

func ToInt(s string) int {
	r, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return r
}

func ToString(s int) string {
	return strconv.Itoa(s)
}

func IsNumeric(s string) bool {
	if IsEmpty(s) {
		return false
	}
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return true
}

func IsAlpha(s string) bool {
	if IsEmpty(s) {
		return false
	}

	runes := []rune(s)
	for _, r := range runes {
		if (r < 65 || r > 122) || (r > 90 && r < 97) {
			return false
		}
	}
	return true
}

func OrElse(origin string, defaultValue string) string {
	if IsEmpty(origin) {
		return defaultValue
	}
	return origin
}
