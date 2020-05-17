# DBML parser for Go

DBML-go is a Go parser for [DBML](https://www.dbml.org) syntax.

## Installation

Go get

```bash
go get github.com/duythinht/dbml-go/...
```

## Quick start

```go
package main

import (
	"fmt"
	"os"

	"github.com/duythinht/dbml-go/parser"
	"github.com/duythinht/dbml-go/scanner"
)

func main() {
	f, _ := os.Open("test.dbml")
	s := scanner.NewScanner(f)
	parser := parser.NewParser(s)
	dbml, err := parser.Parse()
	if err != nil {
        fmt.Printf("%s\n", err)
        os.Exit(1)
    }

    // process dbml here
}
```

