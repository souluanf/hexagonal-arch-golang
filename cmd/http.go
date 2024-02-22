package cmd

import (
	"fmt"
	server2 "github.com/souluanf/hexagonal-arch-golang/adapters/web/server"

	"github.com/spf13/cobra"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "A web server to manage products",
	Long:  `http is a web server to manage products. Use it to get, enable or disable a product, or to create a new product.`,
	Run: func(cmd *cobra.Command, args []string) {
		server := server2.MakeNewWebServer()
		server.Service = &productService
		fmt.Println("Server is running at :8080")
		server.Serve()
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
