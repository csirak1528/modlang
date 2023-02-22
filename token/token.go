package token

type TokenType int

const (
	DOT            TokenType = iota // .
	COMMA                           // ,
	ASSIGN                          // =
	SEMICOLON                       // ;
	COLON                           // :
	LEFT_CURLY                      // {
	RIGHT_CURLY                     // }
	LEFT_PAREN                      // )
	RIGHT_PAREN                     // (
	LEFT_BRACKET                    // [
	RIGHT_BRACKET                   // ]
	BACKWARD_SLASH                  // \
	DOUBLE_QUOTE                    // "
	SINGLE_QUOTE                    // '
	TICK                            // `
	GREATER                         // >
	LESS                            // <
	NOT                             // !

	// MATH
	ADD           // +
	SUB           // -
	STAR          // *
	FORWARD_SLASH // /

	DOLLAR        // $
	QUESTION      // ?
	UNDER_SCORE   // _
	NOT_EQUALS    // !=
	EQUALS        // ==
	LESS_OR_EQ    // <=
	GREATER_OR_EQ // >=
	AND           // &&
	OR            // ||
	EXPONENT      // **
	INC           // ++
	DEC           // --
	ADD_ASSIGN    // +=
	SUB_ASSIGN    // -=
	MUL_ASSIGN    // *=
	DIV_ASSIGN    // /=

	// SPECIAL TOKEN TYPES

	EOF        // end of file
	TYPE       // variable type
	KEYWORD    // keywords like var, module, if
	NUMBER     // number
	IDENTIFIER // identifier
	NULL       // null
	ERROR      // error
	STRING     // string
	START
)

type Token struct {
	Type TokenType
	Data string
}

func CreateAndGetPointer(Type TokenType, Data string) *Token {
	t := Token{Type, Data}
	return &t
}

var TOKENS = []string{"DOT", "COMMA", "ASSIGN", "SEMICOLON", "COLON", "LEFT_CURLY", "RIGHT_CURLY", "LEFT_PAREN", "RIGHT_PAREN", "LEFT_BRACKET", "RIGHT_BRACKET", "FORWARD_SLASH", "BACKWARD_SLASH", "DOUBLE_QUOTE", "SINGLE_QUOTE", "TICK", "GREATER", "LESS", "NOT", "ADD", "SUB", "STAR", "DOLLAR", "QUESTION", "UNDER_SCORE", "NOT_EQUALS", "EQUALS", "LESS_OR_EQ", "GREATER_OR_EQ", "AND", "OR", "EXPONENT", "INC", "DEC", "ADD_ASSIGN", "SUB_ASSIGN", "MUL_ASSIGN", "DIV_ASSIGN", "EOF", "TYPE", "KEYWORD", "NUMBER", "IDENTIFIER", "NULL", "ERROR", "STRING", "START"}

var MATH = []TokenType{ADD, SUB, STAR, FORWARD_SLASH, EXPONENT}

func (t TokenType) ExistsIn(list []TokenType) bool {
	for _, l := range list {
		if t == l {
			return true
		}
	}
	return false
}
