package main

import (
	"fmt"
	"mylang/lexer"
	"mylang/parser"
)

func main() {
	
	lex := lexer.NewLexer(`
	var x = 5 
	var y = 5 
	def name()

	end

	`)
	
	tokens := lex.Lexical_analysis()
	p := parser.NewParser(tokens)
	tree := p.Parse()
	
	printNode(tree, 0)
}

func printNode(n parser.Node, depth int) {
	prefix := ""
	for i := 0; i < depth; i++ {
		prefix += "  "
	}
	fmt.Println(prefix + n.Token.Value)
	for _, child := range n.Children {
		printNode(child, depth+1)
	}
}
