package parser

import (
	"strings"
	"testing"

	"github.com/thanhpd56/dbml-go/scanner"
)

func p(str string) *Parser {
	r := strings.NewReader(str)
	s := scanner.NewScanner(r)
	parser := NewParser(s)
	return parser
}

func TestIllegalSyntax(t *testing.T) {
	parser := p(`Project test { abc , xyz`)
	_, err := parser.Parse()
	if err == nil {
		t.Fail()
	}
}

func TestParseSimple(t *testing.T) {
	parser := p(`
	Project test {
		note: 'just test note'
	}
	table users {
		id int [pk, note: 'just test column note']
	}
	table float_number {
		
	}
	`)
	dbml, err := parser.Parse()
	if err != nil {
		t.Fail()
	}
	if dbml.Project.Name != "test" {
		t.Fail()
	}

	if dbml.Project.Note != "just test note" {
		t.Fail()
	}

	usersTable := dbml.Tables[0]
	if usersTable.Name != "users" {
		t.Fail()
	}
	idColumn := usersTable.Columns[0]
	if idColumn.Name != "id" {
		t.Fail()
	}
	if !idColumn.Settings.PK {
		t.Fail()
	}
	if idColumn.Settings.Note != "just test column note" {
		t.Fail()
	}
}

func TestParseTableName(t *testing.T) {
	parser := p(`
	Table int {
		id int
	}
	`)
	dbml, err := parser.Parse()
	if err != nil {
		t.Fail()
	}
	table := dbml.Tables[0]
	if table.Name != "int" {
		t.Fatalf("table name should be 'int'")
	}
}

func TestParseTableWithType(t *testing.T) {
	parser := p(`
	Table int {
		type int
	}
	`)
	dbml, err := parser.Parse()
	if err != nil {
		t.Fail()
	}
	table := dbml.Tables[0]
	if table.Columns[0].Name != "type" {
		t.Fatalf("column name should be 'type'")
	}
}

func TestParseTableWithNoteColumn(t *testing.T) {
	parser := p(`
	Table int {
		note int
	}
	`)
	dbml, err := parser.Parse()

	//t.Log(err)
	if err != nil {
		t.Fatalf("%v", err)
	}

	table := dbml.Tables[0]
	if table.Columns[0].Name != "note" {
		t.Fatalf("column name should be 'note'")
	}
}
