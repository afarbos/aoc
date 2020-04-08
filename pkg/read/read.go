package read

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func ReadStr(path, sep string) []string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(b), sep)
}

func Read(path, sep string) []int {
	s := ReadStr(path, sep)
	res := make([]int, len(s))
	for index, v := range s {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		res[index] = i
	}
	return res
}
