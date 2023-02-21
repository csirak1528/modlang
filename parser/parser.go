package parser

import (
	"github.com/csirak1528/modlang/src/token"
)

type Parser struct {
	Tokens     []*token.Token
	curToken   int
	OpStack    Stack
	TokenStack Stack
}

func (p *Parser) next() *token.Token {
	p.curToken += 1
	return p.Tokens[p.curToken]
}

func (p *Parser) peek() *token.Token {
	return p.Tokens[p.curToken+1]
}

func (p *Parser) Parse() {
	
}