package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bloom42/libs/rz-go/v2"
	"gitlab.com/bloom42/libs/rz-go/v2/log"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "gon2n",
	Short: "Go bindings and CLI for n2n",
	Long: `Go bindings and CLI for n2n.

Find more information at:
https://pojntfx.github.io/gon2n/`,
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
