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
	Null
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

var Tokens = map[string]int{
	"and":    And,
	"class":  Class,
	"else":   Else,
	"false":  False,
	"fun":    Fun,
	"for":    For,
	"if":     If,
	"null":   Null,
	"or":     Or,
	"print":  Print,
	"return": Return,
	"super":  Super,
	"this":   This,
	"true":   True,
	"var":    Var,
	"while":  While,
}

type Token struct {
	Type    int
	Lexeme  string
	Literal string // JS moment
	Line    int
}

func (t Token) String() string {
	return fmt.Sprintf("[%d] %d %s %s", t.Line, t.Type, t.Lexeme, t.Literal)
}
