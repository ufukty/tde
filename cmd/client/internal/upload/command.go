package upload

import (
	"fmt"
	"tde/internal/command"
)

var DefaultExcludeDirs = []string{".git", "build"}

type Command struct {
	ExcludeDirs command.MultiString `short:"e" long:"exclude-dir"`
}

func (c *Command) Run() {
	c.ExcludeDirs = append(c.ExcludeDirs, DefaultExcludeDirs...)

	for i, dir := range c.ExcludeDirs {
		fmt.Println("i:", i, "dir:", dir)
	}
}
