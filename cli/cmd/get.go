package cmd

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/harryvince/s/internal"
	"github.com/spf13/cobra"
)

var getCommand = &cobra.Command{
	Use:   "get <name>",
	Short: "Gets a secret",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
        env, _ := cmd.Flags().GetString("env")

		config := internal.GetConfig()

		client := internal.NewSSMClient()
		parameter := client.Param(fmt.Sprintf(`/%s/%s/%s`, config.Prefix, env, name))
		value, err := parameter.GetValue()
		if err != nil {
			if awsErr, ok := err.(awserr.Error); ok {
				if awsErr.Code() == "ParameterNotFound" {
					fmt.Println("Unable to find secret, does it exist?")
				} else {
					fmt.Println("AWS Error", awsErr)
				}
			} else {
				fmt.Println("Error updating parameter:", err)
			}
			return
		}
		fmt.Println(value)
	},
}
