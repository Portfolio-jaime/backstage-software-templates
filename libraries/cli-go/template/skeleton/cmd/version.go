package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of ${{values.app_name}}",
	Long:  `All software has versions. This is ${{values.app_name}}'s`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("${{values.app_name}} v1.0.0")
		fmt.Println("Environment: ${{values.app_env}}")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}