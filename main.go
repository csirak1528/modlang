package main

import (
	"fmt"
	"os"

	"github.com/csirak1528/modlang/parser"
	"github.com/csirak1528/modlang/token"

	lex "github.com/csirak1528/modlang/lexer"
)

func main() {
	StackTest()
	// lexer := loadFile()
	// tokens := lexer.Exec()
	// for _, t := range tokens {
	// 	Type := token.TOKENS[t.Type]
	// 	fmt.Println(Type, t.Data)
	// }

	// parser := parser.Parser{Tokens: tokens}
	// parser.Parse()
}

func loadFile() *lex.Lexer {
	fileName := os.Args[1]
	lexer := lex.Lexer{}
	lexer.LoadFromFile(fileName)
	return &lexer
}

func StackTest() {

	tok := &token.Token{Type: token.EOF}
	stack := parser.TokenStackNode{}
	stack.Push(tok)
	stack.Push(tok)
	stack.Push(tok)
	fmt.Println(stack.ToArray())

}
