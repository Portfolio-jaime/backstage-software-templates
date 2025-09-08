package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	version   = "dev"
	commit    = "none"
	date      = "unknown"
	builtBy   = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information",
	Long:  `Print the version information for ${{values.app_name}}`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("${{values.app_name}} version information:\n")
		fmt.Printf("  Version:      %s\n", version)
		fmt.Printf("  Commit:       %s\n", commit)
		fmt.Printf("  Built:        %s\n", date)
		fmt.Printf("  Built by:     %s\n", builtBy)
		fmt.Printf("  Go version:   %s\n", runtime.Version())
		fmt.Printf("  OS/Arch:      %s/%s\n", runtime.GOOS, runtime.GOARCH)
	},
}
