package compiler

type TokenType int
type Token struct {
	Type    TokenType
	Literal string
	Line    uint
}

const (
	TKN_EOF = iota
	TKN_ILLEGAL

	TKN_HASH
	TKN_STRING
)
