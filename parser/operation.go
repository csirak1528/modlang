package parser

import "github.com/csirak1528/modlang/src/token"

type OperationType int

const (
	ASSIGN OperationType = iota // TYPE IDENTIFIER ASSIGN || IDENTIFIER ASSIGN
	MATH
)

type Operation struct {
	Type     int
	Parent   *Operation
	Children []*Operation
	Data     []*token.Token
}

type MathOperation struct {
	Sign        token.TokenType
	Left        *MathOperation
	Right       *MathOperation
	Token       *token.Token
	Operation   *Operation
	IsOperation bool
}

type AssignOperation struct {
	Left  []*token.Token
	Right []*token.Token
}


// NUMBER ADD NUMBER || NUMBER ADD MATH || MATH ADD MATH
// NUMBER SUB NUMBER || NUMBER SUB MATH || MATH SUB MATH
// NUMBER MUL NUMBER || NUMBER MUL MATH || MATH MUL MATH
// NUMBER EXP NUMBER || NUMBER EXP MATH || MATH EXP MATH
// NUMBER FORWARD_SLASH NUMBER || NUMBER FORWARD_SLASH MATH || MATH FORWARD_SLASH MATH
