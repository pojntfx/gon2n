package cmd

import (
	"strings"

	constants "github.com/pojntfx/gon2n/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bloom42/libs/rz-go"
	"gitlab.com/bloom42/libs/rz-go/log"
)

var rootCmd = &cobra.Command{
	Use:   "supernodectl",
	Short: "supernodectl manages supernoded, the n2n supernode management daemon",
	Long: `supernodectl manages supernoded, the n2n supernode management daemon.

Find more information at:
https://pojntfx.github.io/gon2n/`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		viper.SetEnvPrefix("supernode")
		viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	},
}

// Execute starts the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(constants.CouldNotStartRootCommandErrorMessage, rz.Err(err))
	}
}
