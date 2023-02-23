package parser

import (
	"github.com/csirak1528/modlang/token"
)

type Parser struct {
	Tokens    []*token.Token
	curToken  int
	ItemStack *Stack
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
		p.ItemStack.Push(p.getCurToken())
	}
	p.back()
}

func (p *Parser) waitForWithStack(t []token.TokenType, s *Stack) {
	for !p.next().Type.ExistsIn(t) {
		s.Push(p.getCurToken())
	}
	p.back()
}

func (p *Parser) waitForNot(t []token.TokenType) {
	for p.next().Type.ExistsIn(t) {
		p.ItemStack.Push(p.getCurToken())
	}
	p.back()
}
func (p *Parser) ParseAll() *Stack {
	outputStack := &Stack{}
	for {
		curToken := p.getCurToken()
		if curToken == nil || curToken.Type == token.EOF {
			break
		}
		if curToken.Type.ExistsIn(token.MATH) {
			p.ItemStack.Push(p.parseMath())
		} else if curToken.Type == token.LEFT_PAREN {
			p.parseParen()
		} else if curToken.Type == token.NUMBER {
			data := []any{curToken.Data}
			p.ItemStack.Push(&Operation{Type: token.NUMBER, Children: data})
		} else if curToken.Type == token.SEMICOLON {
			outputStack.Push(p.ItemStack.Parse())
		} else {
			p.ItemStack.Push(curToken)
		}
		if p.curToken >= len(p.Tokens)-1 {
			break
		}
		p.next()
	}

	return outputStack
}

func (p *Parser) Parse() *Operation {
	for {
		curToken := p.getCurToken()
		if curToken == nil || curToken.Type == token.EOF {
			break
		}
		if curToken.Type.ExistsIn(token.MATH) {
			p.ItemStack.Push(p.parseMath())
		} else if curToken.Type == token.LEFT_PAREN {
			p.parseParen()
		} else if curToken.Type == token.NUMBER {
			data := []any{curToken.Data}
			p.ItemStack.Push(&Operation{Type: token.NUMBER, Children: data})
		} else if curToken.Type == token.SEMICOLON {
			p.ItemStack.Push(p.ItemStack.Parse())
			p.next()
		} else {
			p.ItemStack.Push(curToken)
		}
		if p.curToken >= len(p.Tokens)-1 {
			break
		}
		p.next()
	}

	out := p.ItemStack.Pop()
	return out.(*Operation)
}

func (p *Parser) parseMath() *Operation {
	operation := p.getCurToken()
	left := p.ItemStack.Parse()
	switch operation.Type {
	case token.LEFT_PAREN:
		p.parseParen()
	case token.EXPONENT:
		p.parseExp()
	case token.STAR:
		p.parseMulDiv()
	case token.FORWARD_SLASH:
		p.parseMulDiv()
	case token.ADD:
		p.parseAddSub()
	case token.SUB:
		p.parseAddSub()
	}

	if p.peek().Type == token.LEFT_PAREN {
		p.next()
		p.parseParen()
	}
	p.ItemStack.Push(token.EOFTOKEN)
	right := p.ItemStack.Parse()
	return CreateMathOperation(operation.Type, left, right)
}

func (p *Parser) parseParen() {
	curStack := &Stack{}
	p.waitForWithStack([]token.TokenType{token.RIGHT_PAREN}, curStack)
	curStack.Push(token.EOFTOKEN)
	p.ItemStack.Push(curStack.Parse())
	p.next()
}

func (p *Parser) parseExp() {
	p.waitForNot([]token.TokenType{token.EXPONENT, token.NUMBER})
}

func (p *Parser) parseMulDiv() {
	p.waitForNot([]token.TokenType{token.STAR, token.FORWARD_SLASH, token.EXPONENT, token.NUMBER})

}

func (p *Parser) parseAddSub() {
	p.waitFor([]token.TokenType{token.EOF, token.SEMICOLON})
}
