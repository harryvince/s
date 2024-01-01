package cmd

import (
	"fmt"

	"github.com/harryvince/s/internal"
	"github.com/spf13/cobra"
)

var updateCommand = &cobra.Command{
	Use:   "update <name> <value>",
	Short: "Updates a secret",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name, value := args[0], args[1]
        env, _ := cmd.Flags().GetString("env")

		config := internal.GetConfig()

		// Write the secret to parameter store
		client := internal.NewSSMClient()
		parameter := client.Param(fmt.Sprintf(`/%s/%s/%s`, config.Prefix, env, name))
		err := parameter.PutValue(value, true)
		if err != nil {
			fmt.Println("Error updating parameter:", err)
			return
		}
		fmt.Println("Updated secret:", name)
	},
}
