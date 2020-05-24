package gen

import (
	"fmt"
	"regexp"
	"strings"
)

type Opts struct {
	From                  string
	Out                   string
	Package               string
	FieldTags             []string
	ShouldGenTblName      bool
	ShouldParseAnnotation bool
	Recursive             bool
	Exclude               string
}

// Generate go model
func Generate(opts Opts) {
	var pattern *regexp.Regexp
	if strings.TrimSpace(opts.Exclude) != "" {
		pattern, _ = regexp.Compile(opts.Exclude)
	}

	dbmls := parseDBML(opts.From, opts.Recursive, opts.ShouldParseAnnotation, pattern)

	g := newgen()
	g.out = opts.Out
	g.gopackage = opts.Package
	g.fieldtags = opts.FieldTags
	g.shouldGenTblName = opts.ShouldGenTblName

	for _, dbml := range dbmls {
		g.dbml = dbml
		if err := g.generate(); err != nil {
			fmt.Printf("Error generate file %s", err)
		}
	}

}
