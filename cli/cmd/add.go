package cmd

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/harryvince/s/internal"
	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use:   "add <name> <value>",
	Short: "Creates a secret",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name, value := args[0], args[1]

		config := internal.GetConfig()

		// Write the secret to parameter store
		client := internal.NewSSMClient()
		parameter := client.Param(fmt.Sprintf(`/%s/%s`, config.Prefix, name))
		err := parameter.PutValue(value, false)
		if err != nil {
			if awsErr, ok := err.(awserr.Error); ok {
				if awsErr.Code() == "ParameterAlreadyExists" {
					fmt.Println("This secret exists to update it use `s update <name> <value>`")
				} else {
					fmt.Println("AWS Error", awsErr)
				}
			} else {
				fmt.Println("Error updating parameter:", err)
			}
			return
		}
		fmt.Println("Created new secret:", name)
	},
}

