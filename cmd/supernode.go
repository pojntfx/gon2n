package cmd

import (
	"github.com/pojntfx/gon2n/pkg"
	"github.com/spf13/cobra"
)

var SupernodeCmd = &cobra.Command{
	Use:   "supernode",
	Short: "Start a n2n supernode",
	Run: func(cmd *cobra.Command, args []string) {
		supernode := pkg.Supernode{}

		supernode.Start()
	},
}

func init() {
	RootCmd.AddCommand(SupernodeCmd)
}
