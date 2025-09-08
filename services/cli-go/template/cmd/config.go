package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration",
	Long:  `Manage configuration for ${{values.app_name}}`,
}

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current configuration",
	Long:  `Display the current configuration values for ${{values.app_name}}`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Current configuration:")
		fmt.Printf("Config file: %s\n", viper.ConfigFileUsed())
		fmt.Printf("Verbose: %v\n", viper.GetBool("verbose"))
		
		// Show all settings
		settings := viper.AllSettings()
		if len(settings) > 0 {
			fmt.Println("\nAll settings:")
			for key, value := range settings {
				fmt.Printf("  %s: %v\n", key, value)
			}
		}
	},
}

var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize configuration file",
	Long:  `Create a default configuration file for ${{values.app_name}}`,
	Run: func(cmd *cobra.Command, args []string) {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error finding home directory: %v\n", err)
			os.Exit(1)
		}

		configPath := fmt.Sprintf("%s/.${{values.app_name}}.yaml", home)
		
		// Check if config file already exists
		if _, err := os.Stat(configPath); err == nil {
			fmt.Printf("Configuration file already exists at %s\n", configPath)
			return
		}

		// Create default config content
		defaultConfig := `# ${{values.app_name}} configuration file
verbose: false

# Add your custom configuration here
# key: value
`

		err = os.WriteFile(configPath, []byte(defaultConfig), 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating config file: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Configuration file created at %s\n", configPath)
	},
}

func init() {
	configCmd.AddCommand(configShowCmd)
	configCmd.AddCommand(configInitCmd)
}