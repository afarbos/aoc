package logging

import (
	"io/ioutil"
	"log"
)

// Disable send log output destination to discard and remove all flags.
func Disable() {
	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)
}

// Flags enable all log flags.
func Flags() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
}
