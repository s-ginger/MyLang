package parser

import (
	"mylang/tokens"
)

type Parser struct {
	Tokens []tokens.Token
	pos int
}

type Node struct {
	Token tokens.Token
	Children []Node
}

func NewParser(tokens []tokens.Token) *Parser {
	return &Parser{Tokens: tokens, pos: 0}
}

func (p *Parser) current() tokens.Token {
	if p.pos >= len(p.Tokens) {
		return tokens.Token{Type: tokens.TokenEOF, Value: ""}
	}
	return p.Tokens[p.pos]
}

func (p *Parser) eat(expected tokens.TokenType) tokens.Token {
	tok := p.current()
	if  tok.Type != expected {
		panic("unexpected token: " + tok.String())
	}
	p.pos++
	return tok
}

func (p *Parser) Parse() Node {
	tok := p.current()

	if tok.Type == tokens.TokenKeyword && tok.Value == "var" {
		return p.parceVarDeclaration()
	}

	panic("unsupported token")
}

func (p *Parser) parceVarDeclaration() Node {
	varTok := p.eat(tokens.TokenKeyword)
	identTok := p.eat(tokens.TokenIdentifier)
	assignTok := p.eat(tokens.TokenAssign)
	valueTok := p.eat(tokens.TokenNumber)

	return Node {
		Token: varTok,
		Children: []Node {
			{ Token: identTok },
			{ Token: assignTok, Children: []Node { { Token: valueTok } } },
		},
	}
}

