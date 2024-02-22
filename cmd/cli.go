package cmd

import (
	"fmt"
	"github.com/souluanf/hexagonal-arch-golang/adapters/cli"

	"github.com/spf13/cobra"
)

var action string
var productId string
var productName string
var productPrice float64

var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "manage products",
	Long:  `cli is a command line tool to manage products. Use it to get, enable or disable a product, or to create a new product.`,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := cli.Run(&productService, action, productId, productName, productPrice)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)
	cliCmd.Flags().StringVarP(&action, "action", "a", "enable", "enable | disable a product")
	cliCmd.Flags().StringVarP(&productId, "id", "i", "", "product id")
	cliCmd.Flags().StringVarP(&productName, "product", "n", "", "product name")
	cliCmd.Flags().Float64VarP(&productPrice, "price", "p", 0.0, "product price")
}
