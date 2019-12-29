package cmd

import (
	"github.com/pojntfx/gon2n/pkg"
	"github.com/spf13/cobra"
)

var EdgeCmd = &cobra.Command{
	Use:   "edge",
	Short: "Start an edge",
	RunE: func(cmd *cobra.Command, args []string) error {
		edge := pkg.Edge{}

		return edge.Start()
	},
}

func init() {
	RootCmd.AddCommand(EdgeCmd)
}
