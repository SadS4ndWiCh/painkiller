package compiler

import (
	"strings"
	"unicode"
)

type Lexer struct {
	src []byte
	len uint

	cur uint
	pek uint
	lin uint
	ch  byte
}

func NewLexer(src []byte) *Lexer {
	return &Lexer{src: src, len: uint(len(src))}
}

func (l *Lexer) readChar() {
	if l.pek >= l.len {
		l.ch = 0
	} else {
		l.ch = l.src[l.pek]
	}

	if l.ch == '\n' {
		l.lin += 1
	}

	l.cur = l.pek
	l.pek += 1
}

func (l *Lexer) skipSpace() {
	for unicode.IsSpace(rune(l.ch)) {
		l.readChar()
	}
}

func (l *Lexer) eatString() string {
	buff := strings.Builder{}

	for l.ch != '\n' && l.ch != 0 {
		buff.WriteByte(l.ch)
		l.readChar()
	}

	return buff.String()
}

func (l *Lexer) Next() Token {
	l.readChar()
	l.skipSpace()

	switch l.ch {
	case '#':
		return Token{TKN_HASH, string(l.ch), l.lin}
	case 0:
		return Token{TKN_EOF, "", l.lin}
	default:
		line := l.lin
		literal := l.eatString()
		return Token{TKN_STRING, literal, line}
	}
}
