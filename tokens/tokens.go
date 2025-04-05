package tokens


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
	TokenReturn     TokenType = "RETURN"
    TokenEOF        TokenType = "EOF"
)


type Token struct {
    Type    TokenType // Тип токена
	Lexeme  string    // Лексема
    Value   string    // Значение 
    Line    int       
    Column  int       
}


