package main

import (
	"fmt"
)

type Scanner struct {
	tokens []Token
}

func (s *Scanner) GetTokens() []Token {
	return s.tokens
}

func (s *Scanner) Parse(source string) {
	fmt.Println(source)
}

func NewScanner() *Scanner {
	return &Scanner{}
}
