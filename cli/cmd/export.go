package cmd

import (
	"fmt"

	"github.com/harryvince/s/internal"
	"github.com/spf13/cobra"
)

var exportCommand = &cobra.Command{
	Use:   "export",
	Short: "Sets all secrets as environment variables",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
        env, _ := cmd.Flags().GetString("env")

		client := internal.NewSSMClient()
		params, err := client.GetAllSecrets(env)
		for _, param := range params {
			command := fmt.Sprintf("export %s=%s", param.Name, param.Value)
			fmt.Println(command)
			if err != nil {
				fmt.Println("Error occured while trying to set:", param.Name)
				return
			}
		}
	},
}
