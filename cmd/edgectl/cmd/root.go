package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bloom42/libs/rz-go"
	"gitlab.com/bloom42/libs/rz-go/log"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "edgectl",
	Short: "edgectl manages edged, the n2n edge management daemon",
	Long: `edgectl manages edged, the n2n edge management daemon.

Find more information at:
https://pojntfx.github.io/gon2n/`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		viper.SetEnvPrefix("edge")
		viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	},
}

// Execute starts the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("Could not start root command", rz.Err(err))
	}
}
