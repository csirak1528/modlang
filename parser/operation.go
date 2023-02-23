package parser

import (
	"fmt"
	"math"
	"strconv"

	"github.com/csirak1528/modlang/token"
)

type Operation struct {
	Type     token.TokenType
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

func CreateOperation(opType token.TokenType, left any, right any) *Operation {
	child := []any{left, right}
	o := Operation{Type: opType, Children: child}
	return &o
}



// (NUMBER || MATH) (ADD || SUB || MUL || EXP || FORWARD_SLASH) (NUMBER || MATH)

type AssignOperation struct {
	Left  []*token.Token
	Right []*token.Token
}

// TYPE IDENTIFIER ASSIGN OPERATION || IDENTIFIER ASSIGN OPERATION

func (o *Operation) Log(indent int) {
	fmt.Print("Operation: " + o.Type.GetLog() + "{")
	for l, item := range o.Children {
		switch item.(type) {
		case *Operation:
			fmt.Println()
			printIndent(indent + 1)
			item.(*Operation).Log(indent + 1)
			printIndent(indent)
		default:
			fmt.Print(item)
			if l< len(o.Children)-1{
				fmt.Print(" ")
			}
		}
	}
	fmt.Println("}")
}

func printIndent(i int) {
	for i > 0 {
		fmt.Print("\t")
		i--
	}
}

func (o *Operation) Eval() int {

	if o.Type == token.NUMBER {
		return o.getNum()
	}

	left := o.Children[0].(*Operation)
	right := o.Children[1].(*Operation)
	switch o.Type {
	case token.STAR:
		return left.Eval() * right.Eval()
	case token.FORWARD_SLASH:
		return left.Eval() / right.Eval()
	case token.ADD:
		return left.Eval() + right.Eval()
	case token.SUB:
		return left.Eval() - right.Eval()
	case token.EXPONENT:
		return int(math.Pow(float64(left.Eval()), float64(right.Eval())))
	}
	return 1

}

func (o *Operation) getNum() int {
	i, e := strconv.Atoi(o.Children[0].(string))
	if e != nil {
		panic(e)
	}
	return i
}
