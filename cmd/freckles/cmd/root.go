package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "freckles",
	Short: "A simple dotfile manager.",
	Example: `  Completion:
    bash:       source <(freckles _carapace)
    elvish:     eval (freckles _carapace|slurp)
    fish:       freckles _carapace | source
    oil:        source <(freckles _carapace)
    nushell:    freckles _carapace | save freckles.nu ; nu -c 'source freckles.nu'
    powershell: freckles _carapace | Out-String | Invoke-Expression
    tcsh:       eval ` + "`" + `freckles _carapace` + "`" + `
    xonsh:      exec($(freckles _carapace))
    zsh:        source <(freckles _carapace)
    `,
	Run: func(cmd *cobra.Command, args []string) {},
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func Execute(version string) error {
	rootCmd.Version = version
	rootCmd.InitDefaultVersionFlag()
	return rootCmd.Execute()
}

func init() {
	rootCmd.InitDefaultHelpFlag()
	carapace.Gen(rootCmd)
}
