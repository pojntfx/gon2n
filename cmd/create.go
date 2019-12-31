package cmd

import (
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create gon2n resources",
}

func init() {
	rootCmd.AddCommand(createCmd)
}
