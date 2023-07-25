package produce

import (
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	// The ".." may change depending on you folder structure
	dir := path.Join(path.Dir(filename), "../../../../examples/word-reverse")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
	spew.Config.Indent = "    "
}

func Test_Binary(t *testing.T) {
	c := Command{
		TestName:   "TDE_WordReverse",
		Population: 10,
		Iterate:    1,
	}
	c.Run()
}
