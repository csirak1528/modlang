package main

import (
	"fmt"
	"os"

	"github.com/csirak1528/modlang/token"

	lex "github.com/csirak1528/modlang/lexer"
)

func main() {
	lexer := loadFile()
	tokens := lexer.Exec()
	for _, t := range tokens {
		Type := token.TOKENS[t.Type]
		fmt.Println(Type, t.Data)
	}
}

func loadFile() *lex.Lexer {
	fileName := os.Args[1]
	lexer := lex.Lexer{}
	lexer.LoadFromFile(fileName)
	return &lexer
}
