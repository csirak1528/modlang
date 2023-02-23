package main

import (
	"fmt"
	"os"

	lex "github.com/csirak1528/modlang/lexer"
	"github.com/csirak1528/modlang/parser"
)

func main() {

	lexer := loadFile()
	tokens := lexer.Exec()

	fmt.Println("\nLexing is complete\n")

	mainParser := parser.Parser{Tokens: tokens, ItemStack: &parser.Stack{}}
	mainParser.ParseAll().Log()
	// stack := mainParser.ParseAll().ToArray().([]any)
	// for lineNumber, out := range stack {
	// 	fmt.Print(lineNumber+1,": ")
	// 	fmt.Println("Result: ", out.(*parser.Operation).Eval())
	// }

}

func loadFile() *lex.Lexer {
	fileName := os.Args[1]
	lexer := lex.Lexer{}
	lexer.LoadFromFile(fileName)
	return &lexer
}
