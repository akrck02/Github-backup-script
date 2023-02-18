package Util

import (
	"strconv"
)

func Int2String(num int) string {
	return strconv.FormatInt(int64(num), 10)
}
