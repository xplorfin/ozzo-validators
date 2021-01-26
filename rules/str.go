package rules

import (
	"strconv"
	"strings"
)

// check wether or not a string has spaces
func IsSpaceless(str string) bool {
	if len(str) == 0 {
		return true
	}
	return !strings.Contains(str, " ")
}

// check wether or not an string is a number
func IsNumeric(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

// check wether or not a string is an int
func IsInt(str string) bool {
	if _, err := strconv.Atoi(str); err == nil {
		return true
	}
	return false
}
