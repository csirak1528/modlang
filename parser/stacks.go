package parser

import "github.com/csirak1528/modlang/src/token"

type Stack interface {
	push()
	pop() *any
}

type TokenStackNode struct {
	cur  *token.Token
	next *token.Token
}

func (tsn *TokenStackNode) push(t *token.Token) {
	tsn.next = tsn.cur
	tsn.cur = t
}

func (tsn *TokenStackNode) pop(t *token.Token) *token.Token {
	out := tsn.cur
	tsn.cur = tsn.next
	return out
}

type OperationStackNode struct {
	cur  *Operation
	next *Operation
}

func (tsn *OperationStackNode) push(t *Operation) {
	tsn.next = tsn.cur
	tsn.cur = t
}

func (tsn *OperationStackNode) pop(t *Operation) *Operation {
	out := tsn.cur
	tsn.cur = tsn.next
	return out
}
