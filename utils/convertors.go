package utils

import (
	"fmt"
	"strconv"
)

func ToString(value any) string {
	return fmt.Sprintf("%v", value)
}

func ToInt(value string) int {
	number, _ := strconv.Atoi(value)
	return number
}
