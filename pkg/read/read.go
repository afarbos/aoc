package read

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const eol = "\n"

// Bytes read a file and return a slice of byte.
func Bytes(path string) []byte {
	log.Println("Reading file", path)
	b, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	return b
}

// String read a file and return a slice of string.
func String(path string) string {
	return string(Bytes(path))
}

// Strings read a file and return a slice of string, splitted with sep.
func Strings(path string, sep ...string) []string {
	var s = eol

	if len(sep) > 0 {
		s = sep[0]
	}

	return strings.Split(String(path), s)
}

// Read read a file and return a slice of int, splitted with sep.
func Read(path string, sep ...string) []int {
	var separator = eol

	if len(sep) > 0 {
		separator = sep[0]
	}

	s := Strings(path, separator)
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

// JSON read a file and load the data in the interface.
func JSON(path string, i interface{}) {
	data := Bytes(path)
	if err := json.Unmarshal(data, i); err != nil {
		log.Fatal(err)
	}
}
