package main

import (
	"encoding/binary"
	"flag"
	"log"

	"github.com/afarbos/aoc/pkg/read"
	"github.com/afarbos/aoc/pkg/utils"
)

const (
	a                = "a"
	z                = "z"
	i                = "i"
	o                = "o"
	l                = "l"
	resNextPassword1 = 1987732088
	resNextPassword2 = 1987732321
)

var flagInput string

func init() {
	utils.Init(&flagInput)
}

func iterPassword(pwd []byte) {
	switch {
	case len(pwd) == 0:
		log.Fatal("No more password")
	case pwd[len(pwd)-1] == z[0]:
		pwd[len(pwd)-1] = a[0]
		iterPassword(pwd[0 : len(pwd)-1])
	default:
		pwd[len(pwd)-1]++
	}
}

func isPasswordValid(pwd []byte) bool {
	overlap, increase := make(map[byte]struct{}), false

	for index, letter := range pwd {
		if letter == i[0] || letter == o[0] || letter == l[0] {
			return false
		}

		if index > 1 {
			if letter == pwd[index-1] && letter != pwd[index-2] {
				overlap[letter] = struct{}{}
			}

			increase = increase || letter == pwd[index-1]+1 && pwd[index-1] == pwd[index-2]+1
		}
	}

	return len(overlap) > 1 && increase
}

func nextPassword(pwd []byte) {
	for {
		iterPassword(pwd)

		if isPasswordValid(pwd) {
			return
		}
	}
}

func main() {
	flag.Parse()

	currentPassword := read.Bytes(flagInput)
	if len(currentPassword) != 0 {
		currentPassword = currentPassword[:len(currentPassword)-1]
	}

	nextPassword(currentPassword)
	utils.AssertEqual(int(binary.BigEndian.Uint32(currentPassword)), resNextPassword1)
	log.Println(string(currentPassword))
	nextPassword(currentPassword)
	utils.AssertEqual(int(binary.BigEndian.Uint32(currentPassword)), resNextPassword2)
	log.Println(string(currentPassword))
}
