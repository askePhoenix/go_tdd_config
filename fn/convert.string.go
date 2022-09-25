package fn

import "strconv"

func IntToString(i int) string {
	return Int64ToString(int64(i))
}

func Int64ToString(i int64)string  {
	return strconv.FormatInt(i, 10)
}
