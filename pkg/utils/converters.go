package utils

import "strconv"

// Int convert string to int return int.
func Int(input string) int {
	result, _ := strconv.Atoi(input)
	return result
}
