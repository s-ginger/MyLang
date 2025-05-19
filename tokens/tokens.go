package tokens



type TokenType string

// типы токенов
const (
    TokenIdentifier TokenType = "IDENTIFIER"
    TokenNumber     TokenType = "NUMBER"
    TokenString     TokenType = "STRING" 

    TokenOperator   TokenType = "OPERATOR" // + - * /
	TokenAssign     TokenType = "ASSIGN" // =
	
    TokenKeyword    TokenType = "KEYWORD" // var 

	TokenLBrace     TokenType = "LBRACE"  // {
	TokenRBrace     TokenType = "RBRACE"  // }

	TokenLParen     TokenType = "LPAREN" // (
	TokenRParen     TokenType = "RPAREN" // )

	TokenReturn     TokenType = "RETURN" // return
    TokenEOF        TokenType = "EOF" // end of file
)


type Token struct {
    Type    TokenType // Тип токена
    Value   string    // Значение     
}

func (t Token) String() string {
    return "("  + string(t.Type) + ", " + t.Value + ") "
}

