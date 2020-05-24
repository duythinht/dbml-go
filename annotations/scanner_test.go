package annotations

import (
	"strings"
	"testing"
)

func sc(s string) *Scanner {
	return NewScanner(strings.NewReader(s))
}

func TestScanEOF(t *testing.T) {
	s := sc("")
	if tok, _ := s.Scan(); tok != EOF {
		t.Fatalf("token should be EOF. Got %v", tok)
	}

	s = sc("   ")
	if tok, _ := s.Scan(); tok != EOF {
		t.Fatalf("token should be EOF. Got %v", tok)
	}

	s = sc("  \n    \t  \n ")
	if tok, _ := s.Scan(); tok != EOF {
		t.Fatalf("token should be EOF. Got %v", tok)
	}
}

func TestScanAnnotation(t *testing.T) {
	s := sc("@gotype")
	tok, lit := s.Scan()
	if tok != ANNOTATION {
		t.Fatalf("token should be ANNOTATION. Got %v", tok)
	}
	if lit != "gotype" {
		t.Fatalf("lit should be gotype(6). Got %s(%d)", lit, len(lit))
	}

	s = sc("@gotype(")
	tok, lit = s.Scan()
	if tok != ANNOTATION {
		t.Fatalf("token should be ANNOTATION. Got %v", tok)
	}
	if lit != "gotype" {
		t.Fatalf("lit should be gotype(6). Got %s(%d)", lit, len(lit))
	}

	s = sc("    @gotype")
	tok, lit = s.Scan()
	if tok != ANNOTATION {
		t.Fatalf("token should be ANNOTATION. Got %v", tok)
	}
	if lit != "gotype" {
		t.Fatalf("lit should be gotype(6). Got %s(%d)", lit, len(lit))
	}

	s = sc("@gotype.!4455")
	tok, _ = s.Scan()
	if tok != ILLEGAL {
		t.Fatalf("token should be ANNOTATION. Got %v", tok)
	}

	s = sc("@got ype")
	tok, lit = s.Scan()
	if tok != ANNOTATION {
		t.Fatalf("token should be ANNOTATION. Got %v", tok)
	}
	if lit != "got" {
		t.Fatalf("lit should be got(3). Got %s(%d)", lit, len(lit))
	}

	s = sc("@@gotype")
	tok, _ = s.Scan()
	if tok != ILLEGAL {
		t.Fatalf("token should be ILLEGAL. Got %v", tok)
	}

	s = sc("@gotype)")
	tok, _ = s.Scan()
	if tok != ILLEGAL {
		t.Fatalf("token should be ILLEGAL. Got %v", tok)
	}
}

func TestScanValue(t *testing.T) {
	s := sc("foo")
	tok, lit := s.Scan()
	if tok != VALUE {
		t.Fatalf("token should be VALUE. Got %v", tok)
	}
	if lit != "foo" {
		t.Fatalf("lit should be foo. Got %s", lit)
	}

	s = sc("  foo  ")
	tok, lit = s.Scan()
	if tok != ILLEGAL {
		t.Fatalf("token should be ILLEGAL. Got %v", tok)
	}

	s = sc("  foo)")
	tok, lit = s.Scan()
	if tok != VALUE {
		t.Fatalf("token should be VALUE. Got %v", tok)
	}
	if lit != "foo" {
		t.Fatalf("lit should be foo. Got %s", lit)
	}

	s = sc("foo@")
	tok, _ = s.Scan()
	if tok != ILLEGAL {
		t.Fatalf("token should be VALUE. Got %v", tok)
	}

	s = sc("foo(")
	tok, _ = s.Scan()
	if tok != ILLEGAL {
		t.Fatalf("token should be ILLEGAL. Got %v", tok)
	}

	s = sc("foo@")
	tok, _ = s.Scan()
	if tok != ILLEGAL {
		t.Fatalf("token should be ILLEGAL. Got %v", tok)
	}
}
