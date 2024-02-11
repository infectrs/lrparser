package lrparser

import (
	"errors"
	"strings"
)

var (
	parseError = errors.New("lrparser: failed parsing")
)

// ParseLr parses the value between a left and a right delimiter from the given input
func ParseLr(input, leftDelimiter, rightDelimiter string) (string, error) {
	startIndex := strings.Index(input, leftDelimiter)

	if startIndex == -1 {
		return "", parseError
	}

	startIndex += len(leftDelimiter)

	endIndex := strings.Index(input[startIndex:], rightDelimiter)

	if endIndex == -1 {
		return "", parseError
	}

	return input[startIndex : startIndex+endIndex], nil
}
