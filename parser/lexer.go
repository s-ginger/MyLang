package parser

import (
	"fmt"
	"regexp"
)

// Тип токена
type TokenType string

// Возможные типы токенов
const (
    TokenIdentifier TokenType = "IDENTIFIER"
    TokenNumber     TokenType = "NUMBER"
    TokenString     TokenType = "STRING"
    TokenOperator   TokenType = "OPERATOR"
    TokenKeyword    TokenType = "KEYWORD"
    TokenEOF        TokenType = "EOF"
)

// Структура токена
type Token struct {
    Type    TokenType // Тип токена
    Value   string    // Значение (например, "123", "if", "+")
    Line    int       // Номер строки
    Column  int       // Номер столбца
}


type Lexer struct {
	//input string
	//pos   int
}

func (l *Lexer) NextToken(text string) []Token { 

	reg, err:= regexp.Compile(`[0-9/*-+]+`)
	if err != nil {
		panic(err)
	}
	res := reg.FindAllString(text, -1)
	
	var result []Token

	for v := range res {
		
		switch res[v] {
		case "+":
			result = append(result, Token{Type: TokenOperator, Value: "+", Line: 0, Column: 0})
		case "-":
			result = append(result, Token{Type: TokenOperator, Value: "-", Line: 0, Column: 0})
		case "*":
			result = append(result, Token{Type: TokenOperator, Value: "*", Line: 0, Column: 0})
		case "/":
			result = append(result, Token{Type: TokenOperator, Value: "/", Line: 0, Column: 0})
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			result = append(result, Token{Type: TokenNumber, Value: res[v], Line: 0, Column: 0})
		default:
			fmt.Println("error")
		}
		
	}
	return result
}

