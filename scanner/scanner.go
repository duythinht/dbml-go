package scanner

import (
	"bufio"
	"bytes"
	"io"

	"github.com/tbobek/dbml-go/token"
)

const eof = rune(0)

// Scanner represents a lexical scanner.
type Scanner struct {
	r  *bufio.Reader
	ch rune // for peek
	l  uint
	c  uint
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	s := &Scanner{r: bufio.NewReader(r), l: 1, c: 0}
	s.next()
	return s
}

// Next return next token and literal value
func (s *Scanner) Read() (tok token.Token, lit string) {
	for isWhitespace(s.ch) {
		s.next()
	}

	// Otherwise read the individual character.
	switch {
	case isLetter(s.ch):
		return s.scanIdent()
	case isDigit(s.ch):
		return s.scanNumber()
	default:
		ch := s.ch
		lit := string(ch)
		s.next()
		switch ch {
		case eof:
			return token.EOF, ""
		case '-':
			return token.SUB, lit
		case '<':
			return token.LSS, lit
		case '>':
			return token.GTR, lit
		case '(':
			return token.LPAREN, lit
		case '[':
			return token.LBRACK, lit
		case '{':
			return token.LBRACE, lit
		case ')':
			return token.RPAREN, lit
		case ']':
			return token.RBRACK, lit
		case '}':
			return token.RBRACE, lit
		case ';':
			return token.SEMICOLON, lit
		case ':':
			return token.COLON, lit
		case ',':
			return token.COMMA, lit
		case '.':
			return token.PERIOD, lit
		case '`':
			return s.scanExpression()
		case '\'', '"':
			return s.scanString(ch)
		case '/':
			if s.ch == '/' {
				return token.COMMENT, s.scanComment()
			}
			return token.ILLEGAL, string(ch)
		}
		return token.ILLEGAL, string(ch)
	}
}

func (s *Scanner) scanComment() string {
	var buf bytes.Buffer
	buf.WriteString("/")
	for s.ch != '\n' && s.ch != eof {
		buf.WriteRune(s.ch)
		s.next()
	}
	return buf.String()
}

func (s *Scanner) scanNumber() (token.Token, string) {
	var buf bytes.Buffer
	countDot := 0
	for isDigit(s.ch) || (s.ch == '.' && countDot < 2) {
		if s.ch == '.' {
			countDot++
		}
		buf.WriteRune(s.ch)
		s.next()
	}
	if countDot < 1 {
		return token.INT, buf.String()
	} else if countDot > 1 {
		return token.ILLEGAL, buf.String()
	}
	return token.FLOAT, buf.String()
}

func (s *Scanner) scanString(quo rune) (token.Token, string) {
	switch quo {
	case '"':
		lit, ok := s.scanTo(quo)
		if ok {
			if token.Lookup(lit) == token.IDENT {
				return token.IDENT, lit
			} else {
				return token.STRING, lit
			}
		}
		return token.ILLEGAL, lit
	case '\'':
		if s.ch != '\'' {
			lit, ok := s.scanTo(quo)
			if ok {
				return token.STRING, lit
			}
			return token.ILLEGAL, lit
		}
		// Handle Triple quote string
		var buf bytes.Buffer
		s.next()
		if s.ch == '\'' { // triple quote string
			s.next()
			count := 0
			for count < 3 {
				switch s.ch {
				case '\'':
					count++
				case eof:
					return token.ILLEGAL, buf.String()
				}
				buf.WriteRune(s.ch)
				s.next()
			}
			return token.TSTRING, buf.String()[:buf.Len()-count]
		}
		return token.ILLEGAL, buf.String()
	default:
		return token.ILLEGAL, string(eof)
	}
}

func (s *Scanner) scanExpression() (token.Token, string) {
	lit, ok := s.scanTo('`')
	if ok {
		return token.EXPR, lit
	}
	return token.ILLEGAL, lit
}

func (s *Scanner) scanTo(stop rune) (string, bool) {
	var buf bytes.Buffer
	for {
		switch s.ch {
		case stop:
			s.next()
			return buf.String(), true
		case '\n', eof:
			return buf.String(), false
		default:
			buf.WriteRune(s.ch)
			s.next()
		}
	}
}

func (s *Scanner) scanIdent() (tok token.Token, lit string) {
	var buf bytes.Buffer
	for {
		buf.WriteRune(s.ch)
		s.next()
		if !isLetter(s.ch) && !isDigit(s.ch) && s.ch != '_' && s.ch != '.' {
			break
		}
	}
	return token.Lookup(buf.String()), buf.String()
}

func (s *Scanner) next() {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		s.ch = eof
		return
	}
	if ch == '\n' {
		s.l++
		s.c = 0
	}
	s.c++
	s.ch = ch
}

// LineInfo return line info
func (s *Scanner) LineInfo() (uint, uint) {
	return s.l, s.c
}
