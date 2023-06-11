package scanner

import (
	"strings"
	"testing"

	"github.com/tbobek/dbml-go/token"
)

func sc(str string) *Scanner {
	return NewScanner(strings.NewReader(str))
}

func TestScanForNumber(t *testing.T) {
	s := sc("123.456")
	if tok, lit := s.Read(); tok != token.FLOAT {
		t.Fatalf("token %s, should be %s, lit %s", tok, token.FLOAT, lit)
	}

	s = sc("123.456i")
	if tok, lit := s.Read(); tok != token.FLOAT {
		t.Fatalf("token %s, should be token.FLOAT, lit %s", tok, lit)
		if tok, lit := s.Read(); tok != token.IDENT {
			t.Fatalf("token %s, should be %s, lit %s", tok, token.IDENT, lit)
		}
	}

	s = sc("123")
	if tok, lit := s.Read(); tok != token.INT {
		t.Fatalf("token %s, should be %s, lit %s", tok, token.INT, lit)
	}

	s = sc("123.2.3")
	if tok, lit := s.Read(); tok != token.ILLEGAL {
		t.Fatalf("token %s, should be %s, lit %s", tok, token.ILLEGAL, lit)
	}
}
