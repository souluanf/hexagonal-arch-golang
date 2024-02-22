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
	Use:   "hexagonal-arch-golang",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
