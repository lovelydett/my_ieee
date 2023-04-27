package utils

import (
	"strconv"
	"strings"
)

func Str_Itoa(i int) string {
	return strconv.Itoa(i)
}

func Str_Concate(strs ...string) string {
	var builder strings.Builder
	for _, str := range strs {
		builder.WriteString(str)
	}
	return builder.String()
}

func Str_Atoi(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}
