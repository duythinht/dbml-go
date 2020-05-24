package annotations

type token int

const (
	ILLEGAL token = iota
	EOF

	// Identifiers
	ANNOTATION
	VALUE

	LPAREN // (
	RPAREN // )
	COMMA  // ,
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",

	ANNOTATION: "ANNOTATION",
	VALUE:      "VALUE",

	LPAREN: "(",
	RPAREN: ")",
	COMMA:  ",",
}
