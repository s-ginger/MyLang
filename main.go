package main

import (
	"fmt"
	"mylang/lexer"
	
)

func main() {
	
	lex := lexer.NewLexer(`
	(defn x 10)
	`)
	
	tokens := lex.Lexical_analysis()
	
	fmt.Println(tokens)
}


