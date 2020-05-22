package gen

import (
	"fmt"
)

type Opts struct {
	From      string
	Out       string
	Package   string
	FieldTags []string
}

// Generate go model
func Generate(ops Opts) error {
	dbml, err := parseDBML(ops.From)
	if err != nil {
		fmt.Printf("Error parse %s", err)
		return err
	}

	g := newgen(dbml, ops.Out, ops.Package)
	return g.generate()
}
