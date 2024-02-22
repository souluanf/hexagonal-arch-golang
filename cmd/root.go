package cmd

import (
	"database/sql"
	dbInfra "github.com/souluanf/hexagonal-arch-golang/adapters/db"
	"github.com/souluanf/hexagonal-arch-golang/application"
	"github.com/spf13/cobra"
	"os"
)

var db, _ = sql.Open("sqlite3", "sqlite3.db")
var productDb = dbInfra.NewProductDb(db)
var productService = application.ProductService{Persistence: productDb}

var rootCmd = &cobra.Command{
	Use:   "hexagonal arch with golang",
	Short: "An example of hexagonal architecture with golang",
	Long:  `This is an example of hexagonal architecture with golang.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
