package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/thanhpd56/dbml-go/internal/gen-go-model/gen"
)

func main() {

	var (
		from             = "database.dbml"
		out              = "model"
		gopackage        = "model"
		fieldtags        = []string{"db", "json", "mapstructure"}
		shouldGenTblName = false
		rememberAlias    = false
		recursive        = false
		exclude          = ""
	)

	cmd := &cobra.Command{
		Use: "dbml-gen-go-model",
		RunE: func(cmd *cobra.Command, args []string) error {
			gen.Generate(gen.Opts{
				From:             from,
				Out:              out,
				Package:          gopackage,
				FieldTags:        fieldtags,
				ShouldGenTblName: shouldGenTblName,
				RememberAlias:    rememberAlias,
				Recursive:        recursive,
				Exclude:          exclude,
			})
			return nil
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
	flags.BoolVarP(&rememberAlias, "remember-alias", "", rememberAlias, "should remember table alias. Only applied if \"from\" is a directory")
	flags.BoolVarP(&recursive, "recursive", "", recursive, "recursive search directory. Only applied if \"from\" is a directory")
	flags.StringVarP(&exclude, "exclude", "E", exclude, "regex for exclude \"from\" files. Only applied if \"from\" is a directory")
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error %s", err)
	}
}
