package gen

import (
	"fmt"
	"regexp"
	"strings"
)

type Opts struct {
	From             string
	Out              string
	Package          string
	FieldTags        []string
	ShouldGenTblName bool
	RememberAlias    bool
	Recursive        bool
	Exclude          string
	ExcludeTables    []string
}

// Generate go model
func Generate(opts Opts) {
	var pattern *regexp.Regexp
	if strings.TrimSpace(opts.Exclude) != "" {
		pattern, _ = regexp.Compile(opts.Exclude)
	}

	dbmls := parseDBML(opts.From, opts.Recursive, pattern)

	g := newgen()
	g.out = opts.Out
	g.gopackage = opts.Package
	g.fieldtags = opts.FieldTags
	g.shouldGenTblName = opts.ShouldGenTblName
	g.excludeTables = opts.ExcludeTables

	for _, dbml := range dbmls {
		g.dbml = dbml
		if err := g.generate(); err != nil {
			fmt.Printf("Error generate file %s", err)
		}
	}

}
