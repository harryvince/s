package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "s",
	Short: "A simple low configuration secrets manager",
	Long: `A secrets manager written in golang to cloudify my secrets
however also teach myself golang in the process. This cli will store
all secrets within aws parameter store for low cost.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.s.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

    // Additional args
    var profileArg string
	initCommand.Flags().StringVar(&profileArg, "profile", "default", "The Profile to use")

    rootCmd.AddCommand(initCommand)
	rootCmd.AddCommand(addCommand)
	rootCmd.AddCommand(updateCommand)
	rootCmd.AddCommand(deleteCommand)
	rootCmd.AddCommand(listCommand)
	rootCmd.AddCommand(getCommand)
	rootCmd.AddCommand(exportCommand)
}
