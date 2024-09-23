package main

import (
	"fmt"
	"strconv"
)

type Scanner struct {
	source  string
	tokens  []Token
	start   int
	current int
	line    int
}

func (s *Scanner) ScanTokens() []Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, Token{
		Type:    EOF,
		Lexeme:  "",
		Literal: "",
		Line:    s.line,
	})
	return s.tokens
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case '(':
		s.addToken(LeftParen)
		break
	case ')':
		s.addToken(RightParen)
		break
	case '{':
		s.addToken(LeftBrace)
		break
	case '}':
		s.addToken(RightBrace)
		break
	case ',':
		s.addToken(Comma)
		break
	case '.':
		s.addToken(Dot)
		break
	case '-':
		s.addToken(Minus)
		break
	case '+':
		s.addToken(Plus)
		break
	case ';':
		s.addToken(Semicolon)
		break
	case '*':
		s.addToken(Star)
		break
	case '!':
		if s.match('=') {
			s.addToken(BangEqual)
		} else {
			s.addToken(Bang)
		}
		break
	case '=':
		if s.match('=') {
			s.addToken(EqualEqual)
		} else {
			s.addToken(Equal)
		}
		break
	case '<':
		if s.match('=') {
			s.addToken(LessEqual)
		} else {
			s.addToken(Less)
		}
		break
	case '>':
		if s.match('=') {
			s.addToken(GreaterEqual)
		} else {
			s.addToken(Greater)
		}
		break
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(Slash)
		}
		break
	case '"':
		for s.peek() != '"' && !s.isAtEnd() {
			if s.peek() == '\n' {
				s.line += 1
			}
			s.advance()
		}

		if s.isAtEnd() {
			fuck(s.line, "Undetermined string")
			return
		}

		s.advance()                                                    // close "
		s.addTokenWithLiteral(String, s.source[s.start+1:s.current-1]) // +- to remove ""

		break
	case ' ':
	case '\t':
	case '\r':
		break
	case '\n':
		s.line += 1
		break
	default:
		if s.isDigit(c) {
			for s.isDigit(s.peek()) {
				s.advance()
			}
			if s.peek() == '.' && s.isDigit(s.peekAfter()) {
				s.advance() // eated .
				for s.isDigit(s.peek()) {
					s.advance()
				}
			}

			number, err := strconv.ParseFloat(s.source[s.start:s.current], 64)
			if err != nil {
				fuck(s.line, "Undetermined int")
			}

			s.addTokenWithLiteral(Number, fmt.Sprintf("%f", number)) // convert back to string, lol
		} else if s.isAlpha(c) {
			for s.isAlpha(s.peek()) {
				s.advance()
			}

			keyword := s.source[s.start:s.current]

			val, ok := Tokens[keyword]
			if !ok {
				val = Identifier
			}

			s.addToken(val)
		} else {
			fuck(s.line, "Unexpected Character")
		}
		break
	}
}

func (s *Scanner) advance() rune {
	c := rune(s.source[s.current])
	s.current += 1
	return c
}

func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return '0'
	}
	return rune(s.source[s.current])
}

func (s *Scanner) peekAfter() rune {
	if s.current+1 >= len(s.source) {
		return '0'
	}
	return rune(s.source[s.current+1])
}

func (s *Scanner) match(c rune) bool {
	if s.isAtEnd() {
		return false
	}
	if rune(s.source[s.current]) != c {
		return false
	}
	s.current += 1
	return true
}

func (s *Scanner) isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func (s *Scanner) isAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func (s *Scanner) isAlphaNumeric(c rune) bool {
	return s.isAlpha(c) || s.isDigit(c)
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) addToken(token int) {
	s.tokens = append(s.tokens, Token{
		Type:    token,
		Lexeme:  s.source[s.start:s.current],
		Literal: "",
		Line:    s.line,
	})
}

func (s *Scanner) addTokenWithLiteral(token int, value string) {
	s.tokens = append(s.tokens, Token{
		Type:    token,
		Lexeme:  s.source[s.start:s.current],
		Literal: value,
		Line:    s.line,
	})
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source:  source,
		tokens:  []Token{},
		start:   0,
		current: 0,
		line:    0,
	}
}
