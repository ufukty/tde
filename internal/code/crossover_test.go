package code

import (
	"fmt"
	"go/ast"
	"os"
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

func TestFindParentNode(t *testing.T) {
	var err error

	c1 := Code{}
	err = c1.LoadFromString(
		`package main
		
		func HelloWorld(a int) int {
			var b int

			if a != 0 {
				if a == 1 {
					return -1
				}
				a = a * 5
				fmt.Println("Hello %s", 
					fmt.Sprintf("world"),
				)
			}

			return a * b
		}`,
	)
	if err != nil {
		t.Error(errors.Wrap(err, "Could not load file"))
	}

	fn1, err := c1.GetFunction("HelloWorld")
	if err != nil {
		t.Error(errors.Wrap(err, "Could not get Function from Code"))
	}

	parentWant := fn1.Root.Body.List[1].(*ast.IfStmt).Body
	child := parentWant.List[1].(*ast.AssignStmt)

	parentGot, childIndex := FindParentNodeAndChildIndex(fn1.Root, child)
	if parentGot.(*ast.BlockStmt) != parentWant {
		t.Errorf("Expected to get '%+v' (addr: %p) (type: %s), got '%+v' (addr: %p) (type: %s) for parent node",
			parentWant, &parentWant, reflect.TypeOf(parentWant), parentGot, &parentGot, reflect.TypeOf(parentGot))
	}
	if childIndex != 1 {
		t.Errorf("Expected to get '1', got '%d' for child node index", childIndex)
	}
}

func TestPickCrossOverPoint(t *testing.T) {
	var err error

	c1 := Code{}
	err = c1.LoadFromString(
		`package main
		
		func HelloWorld(a int) {
			a = a * 5
			if a == 5 {
				return -1
			}
			panic("")
		}`,
	)
	if err != nil {
		t.Error(errors.Wrap(err, "Could not load file"))
	}
	fn1, err := c1.GetFunction("HelloWorld")
	if err != nil {
		t.Error(errors.Wrap(err, "Could not get Function from Code"))
	}

	c2 := Code{}
	err = c2.LoadFromString(
		`package main

		import "fmt"
		
		func HelloWorld(s string) {
			var b int

			if a != 0 {
				if a == 1 {
					return -1
				}
				a = a * 5
				fmt.Println("Hello %s", 
					fmt.Sprintf("world"),
				)
			}

			return a * b
		}`,
	)
	if err != nil {
		t.Error(errors.Wrap(err, "Could not load file"))
	}
	fn2, err := c2.GetFunction("HelloWorld")
	if err != nil {
		t.Error(errors.Wrap(err, "Could not get Function from Code"))
	}

	ok := CrossOver(fn1, fn2)
	if !ok {
		t.Error("CrossOver has not performed")
	}
	ok = CrossOver(fn1, fn2)
	if !ok {
		t.Error("CrossOver has not performed")
	}
	ok = CrossOver(fn1, fn2)
	if !ok {
		t.Error("CrossOver has not performed")
	}
	ok = CrossOver(fn1, fn2)
	if !ok {
		t.Error("CrossOver has not performed")
	}

	fmt.Println("========= fn1 =========")
	c1.Print(os.Stdout)
	fmt.Println("========= fn2 =========")
	c2.Print(os.Stdout)
}
