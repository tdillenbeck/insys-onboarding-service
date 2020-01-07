package str

import (
	"strings"
)

func In(s string, sArr []string) bool {
	for _, ss := range sArr {
		if s == ss {
			return true
		}
	}

	return false
}

func Between(value, beginning, end string) string {
	posFirst := strings.Index(value, beginning)
	if posFirst == -1 {
		return ""
	}

	posLast := strings.Index(value, end)
	if posLast == -1 {
		return ""
	}

	posFirstAdjusted := posFirst + len(beginning)
	if posFirstAdjusted >= posLast {
		return ""
	}

	return value[posFirstAdjusted:posLast]
}

func Before(value string, a string) string {
	pos := strings.Index(value, a)
	if pos == -1 {
		return ""
	}

	return value[0:pos]
}

func After(value string, a string) string {
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}

	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}

	return value[adjustedPos:]
}

func Unique(arr []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range arr {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
