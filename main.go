package main

import (
	"fmt"
	"os"

	"github.com/nzin/calc/parser"
)

func main() {
	expr := ""
	for _, a := range os.Args[1:] {
		expr += a + " "
	}

	p := parser.NewParser(expr)
	node, err := p.Parse()
	if err != nil {
		panic(err)
	}
	value, err := node.Compute()
	if err != nil {
		panic(err)
	}
	fmt.Println(value)
}
