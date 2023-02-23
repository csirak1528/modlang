package parser

import (
	"github.com/csirak1528/modlang/token"
)

type TokenStackNode struct {
	val  *token.Token
	next *TokenStackNode
}
type Stack struct {
	head *TokenStackNode
	Size int
}

func (s *Stack) Push(t *token.Token) {
	cur := TokenStackNode{val: t, next: s.head}
	s.Size++
	s.head = &cur
}

func (s *Stack) Pop() *token.Token {
	out := s.head
	if out != nil {
		s.head = out.next
		s.Size--
		return out.val
	}
	return nil
}

func (s *Stack) Peek() *token.Token {
	return s.head.next.val
}

func (s *Stack) ToArray() []*token.Token {
	arr := make([]*token.Token, s.Size)
	for s.Size > 0 {
		arr[s.Size] = s.Pop()
	}
	return arr
}
