package convert

import "strconv"

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
