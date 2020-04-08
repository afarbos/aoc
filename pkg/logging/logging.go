package logging

import (
	"io/ioutil"
	"log"
)

func Disable() {
	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)
}

func Flags() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
}
