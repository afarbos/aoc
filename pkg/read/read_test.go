package read

import (
	"io/ioutil"
	"log"
	"os"
	"syscall"
	"testing"

	"github.com/afarbos/aoc/pkg/logging"
	"github.com/afarbos/aoc/pkg/test"
)

func createTestFile(t *testing.T, bytes []byte) string {
	f, err := ioutil.TempFile(os.TempDir(), "")
	if err != nil {
		log.Fatal(err)
	}

	test.NoErr(t, ioutil.WriteFile(f.Name(), bytes, 0644))

	return f.Name()
}

func TestRead(t *testing.T) {
	const one string = "1"
	path := createTestFile(t, []byte(one))

	logging.Disable()

	bytes := Bytes(path)
	test.EqualInt(t, len(bytes), 1)
	test.EqualString(t, string(bytes), one)

	str := String(path)
	test.EqualString(t, str, one)

	strs := Strings(path, "")
	test.EqualInt(t, len(strs), 1)
	test.EqualString(t, strs[0], one)

	i := Read(path, "")
	test.EqualInt(t, len(i), 1)
	test.EqualInt(t, i[0], 1)

	var j int

	JSON(path, &j)
	test.EqualInt(t, j, 1)

	test.NoErr(t, syscall.Unlink(path))
}
