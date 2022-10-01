package utils

func GetMinAndMax(params []int64) (min int64, max int64) {
	if len(params) > 0 {
		min = params[0]
		max = params[0]
	}

	for _, value := range params {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}
