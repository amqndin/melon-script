package main

import (
	"fmt"
	"strconv"
	"unicode"
)

const (
	TT_INT    = "TT_INT"
	TT_FLOAT  = "FLOAT"
	TT_PLUS   = "PLUS"
	TT_MINUS  = "MINUS"
	TT_MUL    = "MUL"
	TT_DIV    = "DIV"
	TT_LPAREN = "LPAREN"
	TT_RPAREN = "RPAREN"
)

type Error struct {
	Name    string
	Details string
}

type IllegalCharError struct {
	Error
}

func (e *Error) toString() string {
	return fmt.Sprintf("%s: %s", e.Name, e.Details)
}

type Token struct {
	Type  string
	Value interface{}
}

func NewToken(type_ string, value string) *Token {
	return &Token{Type: type_, Value: value}
}

func (t *Token) String() string {
	if t.Value != "" {
		return fmt.Sprintf("%s:%s", t.Type, t.Value)
	}
	return t.Type
}

type Lexer struct {
	text        string
	currentChar string
	pos         int
}

func (e Lexer) advance() {
	e.pos++
	if e.pos < len(e.text) {
		e.currentChar = string(e.text[e.pos])
	} else {
		e.currentChar = ""
	}
}

func (e Lexer) makeTokens() ([]Token, error) {
	var tokens []Token
	for e.currentChar != "" {
		if e.currentChar == " " || e.currentChar == "\t" {
			e.advance()
		} else if e.currentChar == "+" {
			tokens = append(tokens, *NewToken(TT_PLUS, e.currentChar))
			e.advance()
		} else if e.currentChar == "-" {
			tokens = append(tokens, *NewToken(TT_MINUS, e.currentChar))
			e.advance()
		} else if e.currentChar == "*" {
			tokens = append(tokens, *NewToken(TT_MUL, e.currentChar))
			e.advance()
		} else if e.currentChar == "/" {
			tokens = append(tokens, *NewToken(TT_DIV, e.currentChar))
			e.advance()
		} else if e.currentChar == "(" {
			tokens = append(tokens, *NewToken(TT_LPAREN, e.currentChar))
			e.advance()
		} else if unicode.IsDigit([]rune(e.currentChar)[0]) {
			// e.currentChar.makeNumber()
			tokens = append(tokens, *NewToken(TT_INT, e.currentChar))
			e.advance()
		} else {
		  char := e.currentChar
		  e.advance()
		  return [], IllegalCharError{Error: Error{Name: "Illegal character", Details: char}}
		}
	}
	return tokens
}

func makeNumber(s string) Token {
	numStr := ""
	dotCount := 0

	for _, c := range s {
		if c == '.' {
			if dotCount == 1 {
				break
			}
			dotCount++
			numStr += "."
		} else if unicode.IsDigit(c) || c == '+' || c == '-' {
			numStr += string(c)
		} else {
			break
		}
	}

	if dotCount == 0 {
		val, _ := strconv.Atoi(numStr)
		return Token{"TT_INT", val}
	} else {
		val, _ := strconv.ParseFloat(numStr, 64)
		return Token{"TT_FLOAT", val}
	}
}
