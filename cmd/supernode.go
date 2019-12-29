package cmd

import (
	"github.com/pojntfx/gon2n/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/z0mbie42/rz-go/v2"
	"gitlab.com/z0mbie42/rz-go/v2/log"
)

var supernodeCmd = &cobra.Command{
	Use:   "supernode",
	Short: "Start a supernode",
	RunE: func(cmd *cobra.Command, args []string) error {
		supernode := pkg.Supernode{
			ListenPort:     viper.GetInt(supernodeListenPortKey),
			ManagementPort: viper.GetInt(supernodeManagementPortKey),
		}

		return supernode.Start()
	},
}

func init() {
	var (
		supernodeListenPortFlag     int
		supernodeManagementPortFlag int
	)

	supernodeCmd.PersistentFlags().IntVarP(&supernodeListenPortFlag, supernodeListenPortKey, "l", 1234, "UDP listen port")
	supernodeCmd.PersistentFlags().IntVarP(&supernodeManagementPortFlag, supernodeManagementPortKey, "m", 5645, "UDP management port")

	if err := viper.BindPFlags(supernodeCmd.PersistentFlags()); err != nil {
		log.Fatal(couldNotBindFlagsErrorMessage, rz.Err(err))
	}

	viper.AutomaticEnv()

	rootCmd.AddCommand(supernodeCmd)
}
