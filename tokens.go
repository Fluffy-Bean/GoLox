package main

import (
	"fmt"
)

const (
	// Single Char tokens
	LeftParen = iota
	RightParen
	LeftBrace
	RightBrace
	Comma
	Dot
	Minus
	Plus
	Semicolon
	Slash
	Star

	// One or Two Char tokens
	Bang
	BangEqual
	Equal
	EqualEqual
	Greater
	GreaterEqual
	Less
	LessEqual

	// Literals
	Identifier
	String
	Number

	// Keywords
	And
	Class
	Else
	False
	Fun
	For
	If
	Nah
	Or
	Print
	Return
	Super
	This
	True
	Var
	While

	EOF
)

type Token struct {
	Type    int
	Lexeme  string
	Literal string // JS moment
	Line    int
}

func (t Token) String() string {
	return fmt.Sprintf("[%d] %d %s %s", t.Line, t.Type, t.Lexeme, t.Literal)
}
