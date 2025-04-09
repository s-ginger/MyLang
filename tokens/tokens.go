package tokens

import (
	"fmt"
)

type TokenType string

// типы токенов
const (
    TokenIdentifier TokenType = "IDENTIFIER"
    TokenNumber     TokenType = "NUMBER"
    TokenString     TokenType = "STRING"
    TokenOperator   TokenType = "OPERATOR"
	TokenAssign     TokenType = "ASSIGN"
    TokenKeyword    TokenType = "KEYWORD"
	TokenComment    TokenType = "COMMENT"
	TokenFunction   TokenType = "FUNCTION"
	TokenLParen     TokenType = "LPAREN"
	TokenRParen     TokenType = "RPAREN"
	TokenReturn     TokenType = "RETURN"
    TokenEOF        TokenType = "EOF"
)


type Token struct {
    Type    TokenType // Тип токена
    Value   string    // Значение 
    Line    int       
    Column  int       
}

func (t Token) String() string {
    return "[" + string(t.Type) + "] " + t.Value + " (Line: " + fmt.Sprint(t.Line) + ", Column: " + fmt.Sprint(t.Column) + ")"
}

