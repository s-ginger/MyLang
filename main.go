package main

import (
	"fmt"
	//"os"
	"mylang/lexer"
	//"bufio"

)

func main() {

	/*reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')*/
	
	l := lexer.NewLexer(`
	def add(a, b)
    	return a + b
	end
	add(3, 2)
	`)
	
	tokens := l.Lexical_analysis()
	fmt.Println(tokens)

	
}


