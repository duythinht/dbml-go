package gen

import "fmt"

type Opts struct {
	From             string
	Out              string
	Package          string
	FieldTags        []string
	ShouldGenTblName bool
	RememberAlias    bool
	Recursive        bool
}

// Generate go model
func Generate(opts Opts) {
	dbmls := parseDBML(opts.From, opts.Recursive)

	g := newgen()
	g.out = opts.Out
	g.gopackage = opts.Package
	g.fieldtags = opts.FieldTags
	g.shouldGenTblName = opts.ShouldGenTblName

	for _, dbml := range dbmls {
		g.reset(opts.RememberAlias)
		g.dbml = dbml
		if err := g.generate(); err != nil {
			fmt.Printf("Error generate file %s", err)
		}
	}

}
