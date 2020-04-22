package mathematic

import "math"

func MinInt(nums ...int) int {
	res := math.MaxInt32
	for _, num := range nums {
		if num < res {
			res = num
		}
	}
	return res
}

func MaxInt(nums ...int) int {
	res := math.MinInt32
	for _, num := range nums {
		if num > res {
			res = num
		}
	}
	return res
}

func SumString(f func(string) int, strings ...string) int {
	res := 0

	for _, str := range strings {
		if str == "" {
			continue
		}

		res += f(str)
	}
	return res
}
