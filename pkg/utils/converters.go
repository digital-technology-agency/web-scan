package utils

import "strconv"

func Int(input string) int {
	result, _ := strconv.Atoi(input)
	return result
}
