package tokens


type Token struct {
	Value string
	Type  string
	Line  int
	Col   int
}


const (
	IDENTIFIER = "IDENTIFIER"
	NUMBER = "NUMBER"
	BOOL = "BOOL"
	FLOAT = "FLOAT"
	STRING = "STRING"
	OPERATOR = "OPERATOR"
	IF = "IF"
	ELSE = "ELSE"
	ELSEIF = "ELSEIF"
	WHILE = "WHILE"
	FUNCTION = "FUNCTION"
	FOR = "FOR"
	MATCH = "MATCH"
	EOF = "EOF"
)




