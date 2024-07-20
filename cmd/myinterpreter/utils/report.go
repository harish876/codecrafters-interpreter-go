package utils

import "fmt"

func Report(line int, message string) string {
	return fmt.Sprintf("[line %d] Error: %s\n", line, message)
}
