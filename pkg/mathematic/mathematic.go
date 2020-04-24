package mathematic

import "math"

// MinInt returns the smaller int of nums.
func MinInt(nums ...int) int {
	res := math.MaxInt32
	for _, num := range nums {
		if num < res {
			res = num
		}
	}

	return res
}

// MaxInt returns the larger int of nums.
func MaxInt(nums ...int) int {
	res := math.MinInt32
	for _, num := range nums {
		if num > res {
			res = num
		}
	}

	return res
}

// SumString return the sum of f(string).
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
