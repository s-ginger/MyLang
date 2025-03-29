package parser

import (
	"strconv"
)

type Parser struct {
	Code []Token
}

func NewParser(code []Token) *Parser {
	return &Parser{Code: code}
}

func (p *Parser) Parse() [][]byte {
	var result [][]byte
	for i := 0; i < len(p.Code); i++ {
		switch p.Code[i].Type {
		case TokenNumber:
			result = append(result, []byte{1, byte(StringToInt(p.Code[i].Value))} )
		} 
	}
	return result
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

