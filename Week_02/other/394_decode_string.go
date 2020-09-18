package other

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	n := 2
	stack := make([]string, 0, len(s)/n)

	var tempStr string
	for pos := 0; pos < len(s); pos++ {
		value := s[pos]
		if value >= '0' && value <= '9' {
			stack = append(stack, tempStr)
			tempStr = ""
			stack = append(stack, getNumStr(s, &pos))
			pos--
		} else if value == ']' {
			repTime, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[0 : len(stack)-1]
			tempStr = strings.Repeat(tempStr, repTime)
			tempStr = stack[len(stack)-1] + tempStr
			stack = stack[0 : len(stack)-1]
		} else if value != '[' {
			tempStr = tempStr + string(value)
		}
	}

	var builder strings.Builder
	builder.Grow(len(s) * n)

	for _, str := range stack {
		builder.WriteString(str)
	}

	builder.WriteString(tempStr)

	return builder.String()
}

func getNumStr(s string, pos *int) string {
	var builder strings.Builder
	builder.Grow(3)
	for s[*pos] >= '0' && s[*pos] <= '9' {
		builder.WriteByte(s[*pos])
		*pos++
	}

	return builder.String()
}
