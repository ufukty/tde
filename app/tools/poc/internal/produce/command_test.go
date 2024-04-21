package produce

import (
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func init() {
	err := os.Chdir("../../../../examples/words")
	if err != nil {
		panic(err)
	}
	spew.Config.Indent = "    "
}

func Test_Binary(t *testing.T) {
	c := Command{
		TestName:   "TDE_WordReverse",
		Population: 10,
		Iterate:    10,
	}
	c.Run()
}
