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
	for _, t := range tokens {
		t.Log()
	}

	fmt.Println("\nLexing is complete\n")

	parser := parser.Parser{Tokens: tokens}
	parser.Parse()
}

func loadFile() *lex.Lexer {
	fileName := os.Args[1]
	lexer := lex.Lexer{}
	lexer.LoadFromFile(fileName)
	return &lexer
}
