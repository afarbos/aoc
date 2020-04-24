package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/afarbos/aoc/pkg/read"
	"github.com/afarbos/aoc/pkg/utils"
)

const (
	tenMillion  = 10000000
	fiveZero    = "00000"
	sixZero     = "000000"
	resFiveZero = 346386
	resSixZero  = 9958218
)

var flagInput string

func init() {
	utils.Init(&flagInput)
}

func findHashCoin(secretKey, expectedPrefix string) int {
	for i := 0; i < tenMillion; i++ {
		h := md5.New()
		if _, err := io.WriteString(h, secretKey+strconv.Itoa(i)); err != nil {
			log.Fatal(err)
		}

		if strings.HasPrefix(fmt.Sprintf("%x", h.Sum(nil)), expectedPrefix) {
			return i
		}
	}
	log.Fatal("no hash coin found")

	return 0
}

func main() {
	flag.Parse()

	secretKey := read.String(flagInput)
	if len(secretKey) != 0 {
		secretKey = secretKey[:len(secretKey)-1]
	}

	utils.AssertEqual(findHashCoin(secretKey, fiveZero), resFiveZero)
	utils.AssertEqual(findHashCoin(secretKey, sixZero), resSixZero)
}
