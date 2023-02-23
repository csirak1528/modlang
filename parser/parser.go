package parser

import (
	"github.com/csirak1528/modlang/token"
)

type Parser struct {
	Tokens     []*token.Token
	curToken   int
	TokenStack *Stack
}

func (p *Parser) next() *token.Token {
	p.curToken += 1
	return p.getCurToken()
}

func (p *Parser) back() *token.Token {
	p.curToken -= 1
	return p.getCurToken()
}

func (p *Parser) getCurToken() *token.Token {
	return p.Tokens[p.curToken]
}

func (p *Parser) peek() *token.Token {
	return p.Tokens[p.curToken+1]
}

func (p *Parser) waitFor(t []token.TokenType) {
	for !p.next().Type.ExistsIn(t) {
		p.TokenStack.Push(p.getCurToken())
	}
	p.back()
}

func (p *Parser) Parse() {
	curOperation := &Operation{}
	p.TokenStack = &Stack{}
	for {
		curToken := p.getCurToken()
		if curToken.Type == token.EOF {
			break
		} else if curToken.Type.ExistsIn(token.MATH) {
			mathOp := p.parseMath()
			mathOp.setParent(curOperation)
			mathOp.LogMath()
		} else {
			p.TokenStack.Push(curToken)
		}
		p.next()
	}
}
func (p *Parser) parseMath() *Operation {
	operation := p.getCurToken()
	left := p.TokenStack.ToArray()
	p.waitFor([]token.TokenType{token.EOF, token.SEMICOLON})
	right := p.TokenStack.ToArray()
	return CreateMathOperation(operation.Type, left, right)

}

func (p *Parser) parseParen() {

}
