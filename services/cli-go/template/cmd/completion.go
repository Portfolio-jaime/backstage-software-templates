package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	Long: `Generate completion script for your shell.

To load completions:

Bash:
  $ source <(${{values.app_name}} completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ ${{values.app_name}} completion bash > /etc/bash_completion.d/${{values.app_name}}

  # macOS:
  $ ${{values.app_name}} completion bash > /usr/local/etc/bash_completion.d/${{values.app_name}}

Zsh:
  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ ${{values.app_name}} completion zsh > "${fpath[1]}/_${{values.app_name}}"

  # You will need to start a new shell for this setup to take effect.

fish:
  $ ${{values.app_name}} completion fish | source

  # To load completions for each session, execute once:
  $ ${{values.app_name}} completion fish > ~/.config/fish/completions/${{values.app_name}}.fish

PowerShell:
  PS> ${{values.app_name}} completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> ${{values.app_name}} completion powershell > ${{values.app_name}}.ps1
  # and source this file from your PowerShell profile.
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		default:
			fmt.Fprintf(os.Stderr, "Unsupported shell type %q\n", args[0])
		}
	},
}