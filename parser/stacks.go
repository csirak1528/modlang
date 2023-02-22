package parser

import (
	"fmt"

	"github.com/csirak1528/modlang/token"
)


type TokenStackNode struct {
	val  *token.Token
	next *TokenStackNode
	len  int
}
type Stack struct {
	head any
	Size int
}

func (tsn *TokenStackNode) Push(t *token.Token) {
	cur := TokenStackNode{val: t, next: tsn, len: tsn.len + 1}
	tsn = &cur
}

func (tsn *TokenStackNode) Pop() *token.Token {
	out := tsn.val
	tsn = tsn.next
	return out
}

func (tsn *TokenStackNode) Peek() *token.Token {
	return tsn.next.val
}

func (tsn *TokenStackNode) ToArray() []*token.Token {
	fmt.Print(tsn)
	arr := make([]*token.Token, tsn.len)
	for tsn.len >= 0 {
		arr[tsn.len] = tsn.Pop()
	}
	return arr
}
