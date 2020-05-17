package main

import (
	"fmt"

	"github.com/iancoleman/strcase"
)

func main() {
	for _, str := range []string{
		"hello_world",
		"chaoUrl",
	} {
		fmt.Printf("%s => %s\n", str, strcase.ToCamel(str))
	}

}
