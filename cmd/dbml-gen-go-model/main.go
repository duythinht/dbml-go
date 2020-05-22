package main

import (
	"log"

	"github.com/duythinht/dbml-go/internal/gen-go-model/gen"
	"github.com/spf13/cobra"
)

func main() {

	var (
		from             = "database.dbml"
		out              = "model"
		gopackage        = "model"
		fieldtags        = []string{"db", "json", "mapstructure"}
		shouldGenTblName = false
	)

	cmd := &cobra.Command{
		Use: "dbml-gen-go-model",
		RunE: func(cmd *cobra.Command, args []string) error {
			return gen.Generate(gen.Opts{
				From:             from,
				Out:              out,
				Package:          gopackage,
				FieldTags:        fieldtags,
				ShouldGenTblName: shouldGenTblName,
			})
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	flags := cmd.Flags()
	flags.StringVarP(&from, "from", "f", from, "source of dbml, can be https://dbdiagram.io/... | fire_name.dbml")
	flags.StringVarP(&out, "out", "o", out, "output folder")
	flags.StringVarP(&gopackage, "package", "p", gopackage, "single for multiple files")
	flags.StringArrayVarP(&fieldtags, "fieldtags", "t", fieldtags, "go field tags")
	flags.BoolVarP(&shouldGenTblName, "gen-table-name", "", shouldGenTblName, "should generate \"TableName\" function")
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error %s", err)
	}
}
