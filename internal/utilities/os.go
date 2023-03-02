package utilities

import (
	"fmt"
	"os"
)

func Terminate(msg any) {
	fmt.Println(msg)
	os.Exit(1)
}
