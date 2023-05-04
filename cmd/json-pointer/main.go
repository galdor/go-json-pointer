package main

import (
	"github.com/galdor/go-program"
)

func main() {
	var c *program.Command

	p := program.NewProgram("json-pointer",
		"utilities for the go-json-pointer library")

	c = p.AddCommand("find",
		"extract and print the json value referenced by a pointer", cmdFind)
	c.AddArgument("pointer", "the json pointer")
	c.AddOptionalArgument("path", "the file containing the json document")

	p.ParseCommandLine()
	p.Run()
}

func cmdFind(p *program.Program) {
	// TODO
}
