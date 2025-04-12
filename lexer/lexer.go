package lexer

import (
	tokens "mylang/tokens"
	"regexp"
)


type Lexer struct {
	Source string
}

func NewLexer(source string) *Lexer {
	return &Lexer{Source: source}
}

func (l *Lexer) Lexical_analysis() []tokens.Token {
	numberRegex := regexp.MustCompile(`^\d+$`)
	keywordRegex := regexp.MustCompile(`^var$`)
	identifierRegex := regexp.MustCompile(`\w+`)
	assignRegex := regexp.MustCompile(`^=$`)
	functionRegex := regexp.MustCompile(`^def$`)
	returnRegex := regexp.MustCompile(`^return$`)
	endRegex := regexp.MustCompile(`^end$`)
	lparenRegex := regexp.MustCompile(`^\($`)
	rparenRegex := regexp.MustCompile(`^\)$`)
	operatorRegex := regexp.MustCompile(`==|//|\+|\-|\*|/`)

	re := regexp.MustCompile(`\d+|var|\w+|\S`)
	source := re.FindAllString(l.Source, -1)

	var result []tokens.Token

	for _, v := range source {
		if numberRegex.MatchString(v) {
			result = append(result, tokens.Token{Type: tokens.TokenNumber, Value: v})
		} else if keywordRegex.MatchString(v) {
			result = append(result, tokens.Token{Type: tokens.TokenKeyword, Value: v})
		} else if functionRegex.MatchString(v) {
			result = append(result, tokens.Token{Type: tokens.TokenFunction, Value: v})
		} else if returnRegex.MatchString(v) {
			result = append(result, tokens.Token{Type: tokens.TokenReturn, Value: v})
		} else if endRegex.MatchString(v) {
			result = append(result, tokens.Token{Type: tokens.TokenKeyword, Value: v})
		} else if lparenRegex.MatchString(v) {
			result = append(result, tokens.Token{Type: tokens.TokenLParen, Value: v})
		} else if rparenRegex.MatchString(v) {
			result = append(result, tokens.Token{Type: tokens.TokenRParen, Value: v})
		} else if assignRegex.MatchString(v) {
			result = append(result, tokens.Token{Type: tokens.TokenAssign, Value: v})
		} else if operatorRegex.MatchString(v) {
			result = append(result, tokens.Token{Type: tokens.TokenOperator, Value: v})
		} else if identifierRegex.MatchString(v) {
			result = append(result, tokens.Token{Type: tokens.TokenIdentifier, Value: v})
		}
	}
	return result
}
