package iteration

import "strings"

func Repeat(stringToRpeat string, repeatCount int) string {
	var repeated strings.Builder

	for range repeatCount {
		repeated.WriteString(stringToRpeat)
	}
	return repeated.String()
}
