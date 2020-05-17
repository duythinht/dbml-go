package strcase

import "testing"

func TestIntialismCase(t *testing.T) {
	for in, expect := range map[string]string{
		"get__url":   "get__url",
		"GetUrlOk":   "GetURLOk",
		"chaosUrl":   "chaosURL",
		"Id":         "ID",
		"serve_http": "serve_http",
		"url":        "url",
	} {
		if out := Initialism(in); out != expect {
			t.Fatalf("in: %s, out: %s, expect %s\n", in, out, expect)
		}
	}
}

func TestGetInitialismCamelCase(t *testing.T) {
	for in, expect := range map[string]string{
		"get__url":   "Get_URL",
		"GetUrlOk":   "GetURLOk",
		"chaosUrl":   "ChaosURL",
		"id":         "ID",
		"serve_http": "ServeHTTP",
	} {
		if out := GoInitialismCamelCase(in); out != expect {
			t.Fatalf("in: %s, out: %s, expect %s\n", in, out, expect)
		}
	}

}
