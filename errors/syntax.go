package errors

type SyntaxError string

const TokenNotFoundError = "Error: TokenNotFound, please format correctlty: "
const IncorrectSymbolError = "Error: IncorrectSymbol, please format correctlty: "
const StringSyntaxError = "Error: StringSyntax, please format correctlty: "
const OpSyntaxError = "Error: Syntax, please format correctlty: "
const StackSyntaxError = "Error: Syntax, type of Node is not of Token or Operation: "
