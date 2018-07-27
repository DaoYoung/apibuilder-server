package helpers

import (
	"strconv"
)

func JsonInt(v interface{}) int {
	switch v := v.(type) {
	case int:
		return v
	case float64:
		return int(v)
	case string:
		i, _ := strconv.Atoi(v)
		return i
	default:
		return 0
	}
}