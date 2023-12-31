package cmd

import (
	"fmt"

	"github.com/harryvince/s/cli/pkg"
	"github.com/spf13/cobra"
)

var initCommand = &cobra.Command{
	Use:   "init <prefix> <region>",
	Short: "Creates a secret configuration",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		prefix, region := args[0], args[1]
        profile, _ := cmd.Flags().GetString("profile")
        env, _ := cmd.Flags().GetString("env")

		_, err := internal.NewConfig(prefix, profile, region, env)
		if err != nil {
			fmt.Println("Error occured while trying to create new config:", err)
			return
		}
	},
}
