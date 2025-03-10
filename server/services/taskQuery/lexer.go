// lexer.go - Implements a lexer for tokenising query strings.
package taskquery

import (
	"strings"
	"unicode"
)

type TokenType int

const (
	TokenIdentifier TokenType = iota
	TokenOperator
	TokenString
	TokenNumber
	TokenLParen
	TokenRParen
	TokenLogical
	TokenEOF
)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	input   string
	pos     int
	readPos int
	ch      byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}
	l.pos = l.readPos
	l.readPos++
}

func (l *Lexer) peekChar() byte {
	if l.readPos >= len(l.input) {
		return 0
	}
	return l.input[l.readPos]
}

func (l *Lexer) skipWhitespace() {
	for l.ch != 0 && unicode.IsSpace(rune(l.ch)) {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return unicode.IsLetter(rune(ch)) || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readIdentifier() string {
	pos := l.pos
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.pos]
}

func (l *Lexer) readNumber() string {
	pos := l.pos
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.pos]
}

func (l *Lexer) readString() string {
	l.readChar()
	pos := l.pos
	for l.ch != '"' && l.ch != 0 {
		l.readChar()
	}
	str := l.input[pos:l.pos]
	l.readChar()
	return str
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespace()
	var tok Token

	switch l.ch {
	case 0:
		tok = Token{Type: TokenEOF}
	case '(':
		tok = Token{Type: TokenLParen, Value: string(l.ch)}
		l.readChar()
	case ')':
		tok = Token{Type: TokenRParen, Value: string(l.ch)}
		l.readChar()
	case '"':
		tok.Type = TokenString
		tok.Value = l.readString()
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: TokenOperator, Value: string(ch) + string(l.ch)}
			l.readChar()
		} else {
			tok = Token{Type: TokenOperator, Value: string(l.ch)}
			l.readChar()
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: TokenOperator, Value: string(ch) + string(l.ch)}
			l.readChar()
		} else {
			tok = Token{Type: TokenOperator, Value: string(l.ch)}
			l.readChar()
		}
	case '>', '<':
		ch := l.ch
		if l.peekChar() == '=' {
			l.readChar()
			tok = Token{Type: TokenOperator, Value: string(ch) + string(l.ch)}
			l.readChar()
		} else {
			tok = Token{Type: TokenOperator, Value: string(ch)}
			l.readChar()
		}
	default:
		if isLetter(l.ch) {
			literal := l.readIdentifier()
			upper := strings.ToUpper(literal)
			if upper == "AND" || upper == "OR" {
				tok = Token{Type: TokenLogical, Value: upper}
			} else if upper == "LIKE" {
				tok = Token{Type: TokenOperator, Value: upper}
			} else {
				tok = Token{Type: TokenIdentifier, Value: literal}
			}
			return tok
		} else if isDigit(l.ch) {
			tok.Type = TokenNumber
			tok.Value = l.readNumber()
			return tok
		} else {
			l.readChar()
			return l.NextToken()
		}
	}

	return tok
}
