package parser

// Token represents a lexical token.
type Token int

// TokenInfo stores relevant information about the token during scanning.
type TokenInfo struct {
	Token   Token
	Literal string
}

// TokenLookup is a map, useful for printing readable names of the tokens.
var TokenLookup = map[Token]string{
	OTHER:              "OTHER",
	EOF:                "EOF",
	WS:                 "WS",
	STRING:             "STRING",
	NUMBER:             "NUMBER",
	OPEN_PARENTHESIS:   "(",
	CLOSED_PARENTHESIS: ")",
	PLUS:               "+",
	MINUS:              "-",
	MULTIPLY:           ".",
	DIVIDE:             "/",
}

// String prints a human readable string name for a given token.
func (t Token) String() (print string) {
	return TokenLookup[t]
}

// Declare the tokens here.
const (
	// Special tokens
	// Iota simply starts and integer count
	OTHER Token = iota
	EOF
	WS

	// Main literals
	STRING
	NUMBER

	OPEN_PARENTHESIS
	CLOSED_PARENTHESIS

	// Operators
	PLUS
	MINUS
	MULTIPLY
	DIVIDE
)
