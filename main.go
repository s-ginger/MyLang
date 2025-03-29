package main

import (
	"fmt"
	p "mylang/parser"
)

func main() {
	l := p.Lexer{}
	result := l.NextToken("1 + 2")
	fmt.Println(result)

	parser := p.NewParser(result)
	bytecode := parser.Parse()
	fmt.Println(bytecode)
}


