package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/z0mbie42/rz-go/v2"
	"gitlab.com/z0mbie42/rz-go/v2/log"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "gon2n",
	Short: "Go bindings and CLI for n2n",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		viper.SetEnvPrefix("gon2n")
		viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	},
}

// Execute starts the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("Could not start root command", rz.Err(err))
	}
}
