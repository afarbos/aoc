package convert

import "strconv"

// Atoi convert a slice of string to type int.
func Atoi(strings ...string) ([]int, error) {
	var err error

	res := make([]int, len(strings))

	for index, str := range strings {
		i, e := strconv.Atoi(str)
		if e != nil {
			err = e
		}

		res[index] = i
	}

	return res, err
}

// Btoi convert a bool to type int.
func Btoi(b bool) int {
	if b {
		return 1
	}

	return 0
}
