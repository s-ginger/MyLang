package lexer

import (
	"regexp"
	tokens "mylang/tokens"
)


type Lexer struct {
	Source string
}

func NewLexer(source string) *Lexer {
	return &Lexer{Source: source}
}

func (l *Lexer) Lexical_analysis() []tokens.Token {
	numberRegex := regexp.MustCompile(`^\d+$`)
	keywordRegex := regexp.MustCompile(`^let$`)
	identifierRegex := regexp.MustCompile(`\w+`)

	re := regexp.MustCompile(`\d+|let|\w+|\S`)
	source := re.FindAllString(l.Source, -1)

	var result []tokens.Token

	for _, v := range source {
		
		if numberRegex.MatchString(v) {
			result = append(result, tokens.Token{Type: tokens.TokenNumber, Value: v})
		} else if keywordRegex.MatchString(v) {
			result = append(result, tokens.Token{Type: tokens.TokenAssign, Value: v})
		} else if identifierRegex.MatchString(v) {
			result = append(result, tokens.Token{Type: tokens.TokenIdentifier, Value: v})
		}
	}

	return result
}
