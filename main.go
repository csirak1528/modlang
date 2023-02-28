package main

import (
	"os"

	lex "github.com/csirak1528/modlang/lexer"
	"github.com/csirak1528/modlang/parser"
)

func main() {

	lexer := loadFile()
	tokens := lexer.Exec()
	// fmt.Println(token.LogTokens(tokens))

	mainParser := parser.Parser{Tokens: tokens, ItemStack: &parser.Stack{}}
	mainParser.ParseAll().Log(0)

	// stack := mainParser.ParseAll().ToArray().([]any)
	// for lineNumber, out := range stack {
	// 	fmt.Print(lineNumber+1, ": ")
	// 	fmt.Println("Result: ", out.(*parser.Operation).Eval())
	// }

}

func loadFile() *lex.Lexer {
	fileName := os.Args[1]
	lexer := lex.Lexer{}
	lexer.LoadFromFile(fileName)
	return &lexer
}
