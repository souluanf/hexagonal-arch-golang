package cmd

import (
	"github.com/souluanf/hexagonal-arch-golang/adapters/cli"

	"github.com/spf13/cobra"
)

var action string
var productId string
var productName string
var productPrice float64

var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Run(&productService, action, productId, productName, productPrice)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)
	cliCmd.Flags().StringVarP(&action, "action", "a", "enable", "enable | disable a product")
	cliCmd.Flags().StringVarP(&productId, "id", "i", "", "product id")
	cliCmd.Flags().StringVarP(&productName, "product", "n", "", "product name")
	cliCmd.Flags().Float64VarP(&productPrice, "price", "p", 0.0, "product price")
}
