package main

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
	case "(":
		s.addToken(LeftBrace)
		break
	case ")":
		s.addToken(RightBrace)
		break
	case "{":
		s.addToken(LeftBrace)
		break
	case "}":
		s.addToken(RightBrace)
		break
	case ",":
		s.addToken(Comma)
		break
	case ".":
		s.addToken(Dot)
		break
	case "-":
		s.addToken(Minus)
		break
	case "+":
		s.addToken(Plus)
		break
	case ";":
		s.addToken(Semicolon)
		break
	case "*":
		s.addToken(Star)
		break
	case "!":
		if s.match("=") {
			s.addToken(BangEqual)
		} else {
			s.addToken(Bang)
		}
		break
	case "=":
		if s.match("=") {
			s.addToken(EqualEqual)
		} else {
			s.addToken(Equal)
		}
		break
	case "<":
		if s.match("=") {
			s.addToken(LessEqual)
		} else {
			s.addToken(Less)
		}
		break
	case ">":
		if s.match("=") {
			s.addToken(GreaterEqual)
		} else {
			s.addToken(Greater)
		}
		break
	case "/":
		if s.match("/") {
			for s.peek() != "\n" && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(Slash)
		}
		break
	case " ":
	case "\t":
	case "\r":
		// ignore whitespace
		break
	case "\n":
		s.line += 1
		break
	default:
		fuck(s.line, "Unexpected Char.")
		break
	}
}

func (s *Scanner) advance() string {
	c := string(s.source[s.current])
	s.current += 1
	return c
}

func (s *Scanner) peek() string {
	if s.isAtEnd() {
		return ""
	}
	return string(s.source[s.current])
}

func (s *Scanner) match(expected string) bool {
	if s.isAtEnd() {
		return false
	}
	if string(s.source[s.current]) != expected {
		return false
	}
	s.current += 1
	return true
}

func (s *Scanner) addToken(token int) {
	s.tokens = append(s.tokens, Token{
		Type:    token,
		Lexeme:  s.source[s.start:s.current],
		Literal: "",
		Line:    s.line,
	})
	//fmt.Printf("[%d %d %d] ", s.line, s.current, s.start)
	//fmt.Println(s.tokens[len(s.tokens)-1])
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
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
