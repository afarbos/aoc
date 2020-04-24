package read

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const eol = "\n"

func Bytes(path string) []byte {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return b
}

func String(path string) string {
	return string(Bytes(path))
}

func Strings(path, sep string) []string {
	return strings.Split(String(path), sep)
}

func Read(path, sep string) []int {
	s := Strings(path, sep)
	res := make([]int, len(s))

	for index, v := range s {
		if v == "" {
			continue
		}

		i, err := strconv.Atoi(strings.TrimSuffix(v, eol))
		if err != nil {
			log.Fatal(err)
		}

		res[index] = i
	}

	return res
}
