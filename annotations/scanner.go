package annotations

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

const eof = rune(0)

// Scanner represents a lexical Scanner.
type Scanner struct {
	r  *bufio.Reader
	ch rune
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	s := &Scanner{r: bufio.NewReader(r)}
	return s
}

// NewScanner returns a new instance of Scanner.
func NewStringScanner(s string) *Scanner {
	return NewScanner(strings.NewReader(s))
}

func (s *Scanner) Scan() (token, string) {
	s.next()
	for isWhitespace(s.ch) {
		s.next()
	}

	switch {
	case isAnnotation(s.ch):
		return s.scanAnnotation()
	default:
		switch {
		case s.ch == eof:
			return EOF, ""
		case s.ch == '(':
			return LPAREN, "("
		case s.ch == ')':
			return RPAREN, ")"
		case s.ch == ',':
			return COMMA, ","
		default:
			return s.scanValue()
		}
	}
}

func (s *Scanner) scanAnnotation() (token, string) {
	var buf bytes.Buffer

	// 1st rune to annotation must be a character
	s.next()
	if !isLetter(s.ch) {
		return ILLEGAL, ""
	}
	buf.WriteRune(s.ch)

	// loop until LPAREN
	for {
		s.next()
		if s.ch == ' ' || s.ch == eof {
			break
		}
		if s.ch == '(' {
			s.r.UnreadRune()
			break
		}
		if !isLetter(s.ch) && !isDigit(s.ch) && s.ch != '_' && s.ch != '.' {
			return ILLEGAL, buf.String()
		}
		buf.WriteRune(s.ch)
	}
	return ANNOTATION, buf.String()
}

func (s *Scanner) scanValue() (token, string) {
	var buf bytes.Buffer
	buf.WriteRune(s.ch)

	for {
		s.next()
		if s.ch == eof {
			break
		}
		if s.ch == ')' {
			s.r.UnreadByte()
			break
		}
		if !isLetter(s.ch) && !isDigit(s.ch) && s.ch != '_' && s.ch != '.' && s.ch != '/' && s.ch != ':' {
			return ILLEGAL, buf.String()
		}
		buf.WriteRune(s.ch)
	}

	return VALUE, buf.String()
}

func (s *Scanner) next() {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		s.ch = eof
		return
	}
	s.ch = ch
}
