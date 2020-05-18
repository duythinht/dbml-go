package gen

import (
	"fmt"
)

func Generate(from string, out string, gopackage string) error {
	dbml, err := parseDBML(from)
	if err != nil {
		fmt.Printf("Error parse %s", err)
		return err
	}

	g := newgen(dbml, out, gopackage)
	return g.generate()
}
