package common

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Mins(value int, values ...int) int {
	for _, v := range values {
		if v < value {
			value = v
		}
	}
	return value
}

func Maxs(value int, values ...int) int {
	for _, v := range values {
		if v > value {
			value = v
		}
	}
	return value
}
