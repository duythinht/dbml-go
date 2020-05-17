package main

import (
	"fmt"

	"github.com/duythinht/dbml-go/internal/gen-go-model/strcase"
)

func main() {
	for _, str := range []string{
		"hello_world",
		"chaoUrl",
		"get_url",
	} {
		fmt.Printf("%s => %s\n", str, strcase.GoInitialismCamelCase(str))
		fmt.Printf("%s => %s\n", str, strcase.JSONSnakeCase(str))
	}
}
