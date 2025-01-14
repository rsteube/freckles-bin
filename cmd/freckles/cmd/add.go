package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/rsteube/freckles/pkg/freckles"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [FILE]...",
	Short: "add dotfiles",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			d := freckles.Freckle{Path: arg}
			if err := d.Add(false); err != nil {
				println(err.Error())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	carapace.Gen(addCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			batch := carapace.Batch(
				carapace.ActionFiles(),
			)
			if c.Value == "" {
				batch = append(batch, carapace.ActionCallback(func(c carapace.Context) carapace.Action {
					c.Value = "."
					return carapace.ActionFiles().Invoke(c).ToA()
				}))
			}
			return batch.ToA().ChdirF(traverse.UserHomeDir)
		}),
	)
}
