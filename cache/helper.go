package cache

import (
	"strconv"
	"strings"
)

func isInt(val string) bool {
	_, err := strconv.Atoi(val)
	return err == nil
}

func isBool(s string) bool {
	_, err := strconv.ParseBool(s)
	return err == nil
}

func isFloat(s string) bool {
	if !strings.Contains(s, ".") {
		return false
	}

	_, err := strconv.ParseFloat(s, 32)
	return err == nil
}
