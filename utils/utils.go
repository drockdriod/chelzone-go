package utils

import (
	"strings"
)

func CapitalizeEachWord(params ...string) string {
	titleText := strings.Title(params[0])

	stringSlice := strings.Split(titleText, params[1])

	formattedText := strings.Join(stringSlice, " ")

	return formattedText
}

