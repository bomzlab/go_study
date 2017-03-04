package parsetest

import (
	"strconv"
)

func ToInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func ToString(i int) string {
	return toString(i)
}
