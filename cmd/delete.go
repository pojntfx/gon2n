package cmd

import (
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d"},
	Short:   "Delete gon2n resources",
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
