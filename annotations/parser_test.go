package annotations

import (
	"strings"
	"testing"
)

func p(str string) map[string]string {
	s := NewScanner(strings.NewReader(str))
	return Parse(s)
}

func TestParse(t *testing.T) {
	ret := p("")
	if len(ret) != 0 {
		t.Fatalf("Expect no annotation")
	}

	ret = p("@gotype(github.com/foo/bar:XXX)")
	if ret["gotype"] != "github.com/foo/bar:XXX" {
		t.Fatalf("Expect annotation gotype with value github.com/foo/bar:XXX")
	}

	ret = p("@gotype(github.com/foo/bar:XXX) aaaaa")
	if ret["gotype"] != "github.com/foo/bar:XXX" {
		t.Fatalf("Expect annotation gotype with value github.com/foo/bar:XXX")
	}

	ret = p("@gotype(github.com/foo/bar:XXX), @goflag(foo) ")
	if ret["gotype"] != "github.com/foo/bar:XXX" {
		t.Fatalf("Expect annotation gotype with value github.com/foo/bar:XXX")
	}
	if ret["goflag"] != "foo" {
		t.Fatalf("Expect annotation goflag with value foo")
	}

	ret = p(`
@gotype(github.com/foo/bar:XXX) ,

  @goflag(foo)
  `)
	if ret["gotype"] != "github.com/foo/bar:XXX" {
		t.Fatalf("Expect annotation gotype with value github.com/foo/bar:XXX")
	}
	if ret["goflag"] != "foo" {
		t.Fatalf("Expect annotation goflag with value foo")
	}
}
