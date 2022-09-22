package utils

import "strconv"

func StringToInt64(s string, dval int64) int64 {
	res, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return dval
	}

	return res
}
