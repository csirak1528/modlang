package token

import "fmt"

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
	ADD                             // +
	SUB                             // -
	STAR                            // *
	FORWARD_SLASH                   // /
	DOLLAR                          // $
	QUESTION                        // ?
	UNDER_SCORE                     // _
	NOT_EQUALS                      // !=
	EQUALS                          // ==
	LESS_OR_EQ                      // <=
	GREATER_OR_EQ                   // >=
	AND                             // &&
	OR                              // ||
	EXPONENT                        // **
	INC                             // ++
	DEC                             // --
	ADD_ASSIGN                      // +=
	SUB_ASSIGN                      // -=
	MUL_ASSIGN                      // *=
	DIV_ASSIGN                      // /=
	EOF                             // end of file  // SPECIAL TOKEN TYPES
	TYPE                            // variable type
	KEYWORD                         // keywords like var, module, if
	NUMBER                          // number
	IDENTIFIER                      // identifier
	NULL                            // null
	ERROR                           // error
	STRING                          // string
	FUN                             // function declarations
	SCOPE                           // curly scope
)

type Token struct {
	Type TokenType
	Data string
}

func CreateAndGetPointer(Type TokenType, Data string) *Token {
	t := Token{Type, Data}
	return &t
}

func (t *Token) Log() {
	fmt.Println(TOKENS[t.Type] + ":" + t.Data)
}

func (t *Token) GetLog() string {
	return TOKENS[t.Type] + ":" + t.Data
}

func (t TokenType) Log() {
	fmt.Println(TOKENS[t])
}
func (t TokenType) GetLog() string {
	return TOKENS[t]
}

var TOKENS = []string{
	"DOT",
	"COMMA",
	"ASSIGN",
	"SEMICOLON",
	"COLON",
	"LEFT_CURLY",
	"RIGHT_CURLY",
	"LEFT_PAREN",
	"RIGHT_PAREN",
	"LEFT_BRACKET",
	"RIGHT_BRACKET",
	"BACKWARD_SLASH",
	"DOUBLE_QUOTE",
	"SINGLE_QUOTE",
	"TICK",
	"GREATER",
	"LESS",
	"NOT",
	"ADD",
	"SUB",
	"STAR",
	"FORWARD_SLASH",
	"DOLLAR",
	"QUESTION",
	"UNDER_SCORE",
	"NOT_EQUALS",
	"EQUALS",
	"LESS_OR_EQ",
	"GREATER_OR_EQ",
	"AND",
	"OR",
	"EXPONENT",
	"INC",
	"DEC",
	"ADD_ASSIGN",
	"SUB_ASSIGN",
	"MUL_ASSIGN",
	"DIV_ASSIGN",
	"EOF",
	"TYPE",
	"KEYWORD",
	"NUMBER",
	"IDENTIFIER",
	"NULL",
	"ERROR",
	"STRING",
	"FUN",
	"SCOPE",
}

var MATH = []TokenType{ADD, SUB, STAR, FORWARD_SLASH, EXPONENT}

func (t TokenType) ExistsIn(list []TokenType) bool {
	for _, l := range list {
		if t == l {
			return true
		}
	}
	return false
}

func LogTokens(tkns []*Token) string {
	out := tkns[0].GetLog()
	for _, t := range tkns[1:] {
		out += ", " + t.GetLog()
	}
	return out
}

var EOFTOKEN = &Token{Type: EOF}
