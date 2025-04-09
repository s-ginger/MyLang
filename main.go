package main

import (
	"fmt"

	"mylang/lexer"
	//"mylang/tokens"
)

func main() {

	l := lexer.NewLexer(`
	var a = 5
	var a = 8
	a == 2
	a / q
	`)

	tokens := l.Lexical_analysis()
	fmt.Println(tokens)

	fmt.Scan()
}


