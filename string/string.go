package string

import (
	"strconv"
)

//IsEmpty - check is empty string
func IsEmpty(s string) bool {
	return len(s) == 0
}

//ToInt - cast string to int
func ToInt(s string) int {
	r, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return r
}

//ToString - cast int to string
func ToString(s int) string {
	return strconv.Itoa(s)
}

//IsNumeric - check is numeric string
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

//IsAlpha - check is alpha string
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
