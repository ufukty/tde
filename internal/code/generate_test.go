package code

import (
	"fmt"
	"os"
	"testing"
)

func TestPickRandomNodeType(t *testing.T) {
	fmt.Println("orderedNodeTypes", orderedNodeTypes)
	fmt.Println("cumulativeProbabilities", cumulativeProbabilities)

	freq := map[NodeType]int{}
	for i := 0; i < 100000; i++ {
		randomNode := PickRandomNodeType()
		if _, ok := freq[randomNode]; !ok {
			freq[randomNode] = 0
		}
		freq[randomNode]++
	}

	for nodeType, freq_ := range freq {
		fmt.Printf("%-20s: %d\n", nodeType, freq_)
	}
}

func TestNewLine(t *testing.T) {
	txt := `package main
	
	import "fmt"

	func Main() {
		return
	}
	`

	c := Code{}
	c.LoadFromString(txt)
	f, _ := c.FindFunction("Main")

	NewLine(f)

	c.Print(os.Stdout)
}
