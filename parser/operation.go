package parser

import (
	"github.com/csirak1528/modlang/token"
)

type Operation struct {
	Type     token.TokenType
	Parent   *Operation
	Children []any
}

type MathOperation struct {
	Sign        token.TokenType
	Left        *MathOperation
	Right       *MathOperation
	Token       *token.Token
	Operation   *Operation
	IsOperation bool
}

func CreateMathOperation(opType token.TokenType, left any, right any) *Operation {
	child := []any{left, right}
	o := Operation{Type: opType, Children: child}
	return &o
}

func (o *Operation) setParent(p *Operation) {
	o.Parent = p
}

func (o *Operation) LogMath() {
	o.Type.Log()
	for _, c := range o.Children.([]*token.Token) {
		c[0].Log()
	}
}

// (NUMBER || MATH) (ADD || SUB || MUL || EXP || FORWARD_SLASH) (NUMBER || MATH)

type AssignOperation struct {
	Left  []*token.Token
	Right []*token.Token
}

// TYPE IDENTIFIER ASSIGN OPERATION || IDENTIFIER ASSIGN OPERATION
