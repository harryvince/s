package cmd

import (
	"fmt"

	"github.com/harryvince/s/internal"
	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "Lists all secrets",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		ssm := internal.NewSSMClient()
		params, err := ssm.ListSecretNames()
		if err != nil {
			fmt.Println("Error occured while listing secrets:", err)
			return
		}

		fmt.Println("Configured Secrets:", params)
	},
}