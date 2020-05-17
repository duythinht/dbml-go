package token

import "strings"

//Token
type Token int

//go:generate stringer -type=Token
const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	COMMENT

	_literalBeg
	// Identifiers and basic type literals
	// (these tokens stand for classes of literals)
	IDENT   // main
	INT     // 12345
	FLOAT   // 123.45
	IMAG    // 123.45i
	STRING  // 'abc'
	DSTRING // "abc"
	TSTRING // '''abc'''

	EXPR // `now()`

	_literalEnd

	_operatorBeg

	SUB // -
	LSS // <
	GTR // >

	LPAREN // (
	LBRACK // [
	LBRACE // {
	COMMA  // ,
	PERIOD // .

	RPAREN    // )
	RBRACK    // ]
	RBRACE    // }
	SEMICOLON // ;
	COLON     // :

	_operatorEnd

	_keywordBeg

	PROJECT
	TABLE
	ENUM
	REF
	AS
	TABLEGROUP

	_keywordEnd

	_miscBeg

	PRIMARY
	KEY
	PK
	NOTE
	UNIQUE
	NOT
	NULL
	INCREMENT
	DEFAULT

	INDEXES
	TYPE
	DELETE
	UPDATE
	NO
	ACTION
	RESTRICT
	SET

	_miscEnd
)

// Tokens map to string
var Tokens = [...]string{
	ILLEGAL: "ILLEGAL",

	EOF:     "EOF",
	COMMENT: "COMMENT",

	IDENT:   "IDENT",
	INT:     "INT",
	FLOAT:   "FLOAT",
	IMAG:    "IMAG",
	STRING:  "STRING",
	DSTRING: "DSTRING",
	TSTRING: "TSTRING",
	EXPR:    "EXPR",

	SUB: "-",
	LSS: "<",
	GTR: ">",

	LPAREN: "(",
	LBRACK: "[",
	LBRACE: "{",

	RPAREN: ")",
	RBRACK: "]",
	RBRACE: "}",

	SEMICOLON: ";",
	COLON:     ":",
	COMMA:     ",",
	PERIOD:    ".",

	PROJECT:    "PROJECT",
	TABLE:      "TABLE",
	ENUM:       "ENUM",
	REF:        "REF",
	AS:         "AS",
	TABLEGROUP: "TABLEGROUP",

	PRIMARY:   "PRIMARY",
	KEY:       "KEY",
	PK:        "PK",
	NOTE:      "NOTE",
	UNIQUE:    "UNIQUE",
	NOT:       "NOT",
	NULL:      "NULL",
	INCREMENT: "INCREMENT",
	DEFAULT:   "DEFAULT",

	INDEXES:  "INDEXES",
	TYPE:     "TYPE",
	DELETE:   "DELETE",
	UPDATE:   "UPDATE",
	NO:       "NO",
	ACTION:   "ACTION",
	RESTRICT: "RESTRICT",
	SET:      "SET",
}

var keywords map[string]Token

func init() {
	keywords = make(map[string]Token)
	for i := _keywordBeg + 1; i < _miscEnd; i++ {
		keywords[Tokens[i]] = i
	}
}

// Lookup find token with a name
func Lookup(ident string) Token {
	if tok, ok := keywords[strings.ToUpper(ident)]; ok {
		return tok
	}
	return IDENT
}
