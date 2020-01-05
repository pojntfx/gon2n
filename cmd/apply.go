package cmd

import (
	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use:     "apply",
	Aliases: []string{"a"},
	Short:   "Apply gon2n resources",
}

func init() {
	rootCmd.AddCommand(applyCmd)
}
