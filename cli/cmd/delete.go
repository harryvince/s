package cmd

import (
	"fmt"

	"github.com/harryvince/s/internal"
	"github.com/spf13/cobra"
)

var deleteCommand = &cobra.Command{
	Use:   "delete <name>",
	Short: "Deletes a secret",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		env, _ := cmd.Flags().GetString("env")

		config := internal.GetConfig()

		// Write the secret to parameter store
		client := internal.NewSSMClient()
		parameter := client.Param(fmt.Sprintf(`/%s/%s/%s`, config.Prefix, env, name))
		err := parameter.DeleteValue()
		if err != nil {
			fmt.Println("Error deleting parameter:", err)
			return
		}
		fmt.Println("Deleted secret:", name)
	},
}
