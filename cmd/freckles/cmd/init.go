package cmd

import (
	"os"
	"os/exec"

	"github.com/carapace-sh/carapace"
	spec "github.com/carapace-sh/carapace-spec"
	"github.com/rsteube/freckles/pkg/freckles"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init freckles folder",
	RunE: func(cmd *cobra.Command, args []string) error {
		// ANCHOR: command
		c := exec.Command("git", "init", freckles.FreckleDir())
		if cmd.Flag("clone").Changed {
			c = exec.Command("git", "clone", cmd.Flag("clone").Value.String(), freckles.FreckleDir())
		}
		// ANCHOR_END: command
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		if err := c.Run(); err != nil {
			return err
		}

		if _, err := os.Stat(freckles.FreckleDir() + ".frecklesignore"); os.IsNotExist(err) {
			return os.WriteFile(freckles.FreckleDir()+".frecklesignore", []byte(".git\n.frecklesignore\n"), os.ModePerm)
		}
		return nil
	},
}

func init() {
	initCmd.Flags().String("clone", "", "clone existing repo")

	rootCmd.AddCommand(initCmd)

	// ANCHOR: flagcompletion
	carapace.Gen(initCmd).FlagCompletion(carapace.ActionMap{
		"clone": spec.ActionMacro("$carapace.tools.git.RepositorySearch"),
	})
	// ANCHOR_END: flagcompletion
}
