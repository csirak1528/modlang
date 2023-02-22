package lexer

import (
	"io/ioutil"
	"log"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/csirak1528/modlang/errors"
	"github.com/csirak1528/modlang/token"
)

type Lexer struct {
	data  string
	start int
	pos   int
	width int
}

// TODO
// - Add support for addresses
// - Add support for bytes

const SYMBOLS = ".,=;:{})([]/\\\"'`!===<><=>=&&!||+-***?$_"
const WHITESPACE = " \n\t"
const ALPHANUMERIC = "1234567890abcdefghjiklmnopqrstuvwxyzABCDEFGHJIKLMNOPQRSTUVWXYZ"

const STRINGCONTENTS = "1234567890abcdefghjiklmnopqrstuvwxyzABCDEFGHJIKLMNOPQRSTUVWXYZ.,=;:{})([]/'`!===<><=>=&&!||+-***?$_\a\b\f\n\r\t\v'\\ "
const escapeChars = "abfnrtv'\"\\"

var KEYWORDS = []string{"module", "if", "else", "for", "while", "break", "wallet", "this", "const", "import", "return", "struct", "null", "new", "from", "require"}
var TYPES = []string{"uint", "address", "bool", "string", "bytes"}
var DOUBLESYMBOLS = []string{"!", "=", "<", ">", "&", "|", "*"}

type StringArr []string

func existsIn(s []string, key string) bool {
	for _, k := range s {
		if key == k {
			return true
		}
	}
	return false
}
func existsInRune(s string, key rune) bool {
	for _, k := range s {
		if key == k {
			return true
		}
	}
	return false
}

func isAlphaNumeric(r rune) bool {
	if !(unicode.IsLetter(r) || unicode.IsDigit(r)) {
		return false
	}
	return true
}

// Load file in

func (l *Lexer) LoadFromFile(filename string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)
	l.data = text
	l.start = 0
	l.pos = 0
}

// Helpers

func (l *Lexer) next() rune {
	if l.pos >= len(l.data) {
		return -1
	}
	r, width := utf8.DecodeRuneInString(l.data[l.pos:])
	l.width = width
	l.pos += width
	return r
}

func (l *Lexer) backUp() {
	l.pos -= l.width
	l.width = 0
}

func (l *Lexer) peek() rune {
	val := l.next()
	l.backUp()
	return val
}

func (l *Lexer) ignore() {
	l.start = l.pos
	l.width = 0
}

func (l *Lexer) getCurrent() string {
	return l.data[l.start:l.pos]
}

func (l *Lexer) accept(value string) bool {
	val := l.next()
	if val >= 0 && strings.IndexRune(value, val) >= 0 {
		return true
	}
	l.backUp()
	return false
}

func (l *Lexer) acceptRun(value string) {
	for l.accept(value) {
	}
}

// Main Loop

func (l *Lexer) Exec() []*token.Token {
	var tokens []*token.Token
	for {
		l.ignore()
		var curToken *token.Token
		curToken = l.GetNextToken()
		tokens = append(tokens, curToken)
		if curToken.Type == token.EOF {
			break
		}
	}
	return tokens
}

func (l *Lexer) GetNextToken() *token.Token {
	char := l.peek()
	if char == -1 {
		return token.CreateAndGetPointer(token.EOF, "")
	}
	if existsInRune(WHITESPACE, char) {
		l.next()
		l.ignore()
		return l.GetNextToken()
	}
	if unicode.IsDigit(char) {
		return l.lexNumber()
	}
	if isAlphaNumeric(char) {
		return l.lexAlphaNumeric()
	}
	if existsInRune(SYMBOLS, char) {
		return l.lexSymbol()
	}
	panic(errors.TokenNotFoundError + string(char))
}

func (l *Lexer) lexAlphaNumeric() *token.Token {
	l.acceptRun(ALPHANUMERIC)

	term := l.getCurrent()
	if existsIn(KEYWORDS, term) {
		return token.CreateAndGetPointer(token.KEYWORD, term)
	}
	if existsIn(TYPES, term) {
		return token.CreateAndGetPointer(token.TYPE, term)
	}
	return token.CreateAndGetPointer(token.IDENTIFIER, term)
}

func (l *Lexer) lexSymbol() *token.Token {
	var symbolType token.TokenType

	symbol := string(l.next())
	if symbol == "\""  {
		return l.lexString()
	}
	if symbol == "-" && unicode.IsDigit(l.peek()) {
		return l.lexNumber()
	}
	if existsIn(DOUBLESYMBOLS, symbol) && existsInRune(SYMBOLS, l.peek()) {
		l.next()
	}

	switch string(symbol) {
	case ".":
		symbolType = token.DOT
	case ",":
		symbolType = token.COMMA
	case "=":
		symbolType = token.ASSIGN
	case ";":
		symbolType = token.SEMICOLON
	case ":":
		symbolType = token.COLON
	case "{":
		symbolType = token.LEFT_CURLY
	case "}":
		symbolType = token.RIGHT_CURLY
	case ")":
		symbolType = token.LEFT_PAREN
	case "(":
		symbolType = token.RIGHT_PAREN
	case "[":
		symbolType = token.LEFT_BRACKET
	case "]":
		symbolType = token.RIGHT_BRACKET
	case "/":
		symbolType = token.FORWARD_SLASH
	case "\\":
		symbolType = token.BACKWARD_SLASH
	case "\"":
		symbolType = token.DOUBLE_QUOTE
	case "'":
		symbolType = token.SINGLE_QUOTE
	case "`":
		symbolType = token.TICK
	case ">":
		symbolType = token.GREATER
	case "<":
		symbolType = token.LESS
	case "!":
		symbolType = token.NOT
	case "+":
		symbolType = token.ADD
	case "-":
		symbolType = token.SUB
	case "_":
		symbolType = token.UNDER_SCORE
	case "*":
		symbolType = token.STAR
	case "$":
		symbolType = token.DOLLAR
	case "?":
		symbolType = token.QUESTION
	case "!=":
		symbolType = token.NOT_EQUALS
	case "==":
		symbolType = token.EQUALS
	case "<=":
		symbolType = token.LESS_OR_EQ
	case ">=":
		symbolType = token.GREATER_OR_EQ
	case "&&":
		symbolType = token.AND
	case "||":
		symbolType = token.OR
	case "**":
		symbolType = token.EXPONENT
	default:
		panic(errors.IncorrectSymbolError + symbol)
	}
	return token.CreateAndGetPointer(symbolType, "")
}

func (l *Lexer) lexNumber() *token.Token {
	l.accept("+-")
	digits := "1234567890"

	if l.accept("0") && l.accept("xX") {
		digits = "1234567890abcdefABCDEF"
	}

	l.acceptRun(digits)

	if l.accept(".") {
		l.acceptRun(digits)
	}
	return token.CreateAndGetPointer(token.NUMBER, l.getCurrent())
}

func (l *Lexer) lexString() *token.Token {
	l.accept("\"")
	l.acceptRun(STRINGCONTENTS)
	if !l.accept("\"") {
		panic(errors.StringSyntaxError + l.getCurrent() + string(l.peek()))
	}
	return token.CreateAndGetPointer(token.STRING, l.getCurrent())

}
