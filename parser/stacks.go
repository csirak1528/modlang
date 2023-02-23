package parser

import (
	"fmt"

	"github.com/csirak1528/modlang/errors"

	"github.com/csirak1528/modlang/token"
)

type StackNode struct {
	val  any
	next *StackNode
}
type Stack struct {
	head *StackNode
	Size int
}

func (s *Stack) Push(t any) {
	cur := StackNode{val: t, next: s.head}
	s.Size++
	s.head = &cur
}

func (s *Stack) Pop() any {
	out := s.head
	if out != nil {
		s.head = out.next
		s.Size--
		return out.val
	}
	return nil
}

func (s *Stack) Peek() any {
	return s.head.next.val
}

func (s *Stack) Cur() any {
	return s.head.val
}


func (s *Stack) ToArray() any {
	arr := make([]any, s.Size)
	for s.Size > 0 {
		arr[s.Size] = s.Pop()
	}
	return arr
}

func (s *Stack) Parse() *Operation {
	stackItems := s.ToArray().([]any)
	tokens := make([]*token.Token, len(stackItems))
	opStack := &Stack{}
	for i, item := range stackItems {
		switch item.(type) {
		case *token.Token:
			tokens[i] = item.(*token.Token)
		case *Operation:
			opStack.Push(item.(*Operation))
		default:
			panic(errors.StackSyntaxError)
		}
	}

	parser := Parser{Tokens: tokens, ItemStack: opStack}
	return parser.Parse()
}

func (s *Stack) Log() {
	arr := s.ToArray().([]any)
	for _, item := range arr {
		switch item.(type) {
		case *token.Token:
			fmt.Print(item.(*token.Token).GetLog() + " ")
		case *Operation:
			item.(*Operation).Log(0)
		}
	}
	for _, item := range arr {
		s.Push(item)
	}
	fmt.Println()
}
