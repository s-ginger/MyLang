package main

import (
	"fmt"
	"mylang/lexer"
	
)

func main() {
	
	lex := lexer.NewLexer(`
	def add(a, b) {
    	return a + b
	}
	add(3, 2)
	`)
	
	tokens := lex.Lexical_analysis()
	
	fmt.Println(tokens)
}


