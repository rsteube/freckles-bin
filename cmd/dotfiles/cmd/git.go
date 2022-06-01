package cmd

import (
	"os"
	"os/exec"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/bridge"
	"github.com/rsteube/dotfiles-bin/pkg/dotfiles"
	"github.com/spf13/cobra"
)

var gitCmd = &cobra.Command{
	Use:                "git",
	Short:              "invoke git on dotfile directory",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("git", append([]string{"-C", dotfiles.DotfileDir()}, args...)...)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
	},
}

func init() {
	rootCmd.AddCommand(gitCmd)

	carapace.Gen(gitCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return bridge.ActionCarapaceBin("git").Chdir(dotfiles.DotfileDir())
		}),
	)
}
