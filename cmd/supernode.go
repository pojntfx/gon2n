package cmd

import (
	"github.com/pojntfx/gon2n/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/z0mbie42/rz-go/v2"
	"gitlab.com/z0mbie42/rz-go/v2/log"
	"os"
	"os/signal"
	"syscall"
)

var supernodeCmd = &cobra.Command{
	Use:   "supernode",
	Short: "Start a supernode",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !(viper.GetString(supernodeConfigFileKey) == supernodeConfigFileDefault) {
			viper.SetConfigFile(viper.GetString(supernodeConfigFileKey))

			if err := viper.ReadInConfig(); err != nil {
				return err
			}
		}

		supernode := pkg.Supernode{
			ListenPort:     viper.GetInt(supernodeListenPortKey),
			ManagementPort: viper.GetInt(supernodeManagementPortKey),
		}

		if err := supernode.Configure(); err != nil {
			return err
		}

		if err := supernode.OpenListenPortSocket(); err != nil {
			return err
		}

		if err := supernode.OpenManagementPortSocket(); err != nil {
			return err
		}

		interrupt := make(chan os.Signal, 2)
		signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-interrupt

			// Allow manually killing the process
			go func() {
				<-interrupt

				os.Exit(1)
			}()

			log.Info("Gracefully stopping supernode (this might take a few seconds)")

			if err := supernode.Stop(); err != nil {
				log.Fatal("Could not stop supernode", rz.Err(err))
			}
		}()

		log.Info("Starting supernode")

		return supernode.Start()
	},
}

func init() {
	var (
		supernodeConfigFileFlag     string
		supernodeListenPortFlag     int
		supernodeManagementPortFlag int
	)

	supernodeCmd.PersistentFlags().StringVarP(&supernodeConfigFileFlag, supernodeConfigFileKey, "f", supernodeConfigFileDefault, "Configuration file to use")
	supernodeCmd.PersistentFlags().IntVarP(&supernodeListenPortFlag, supernodeListenPortKey, "l", 1234, "UDP listen port")
	supernodeCmd.PersistentFlags().IntVarP(&supernodeManagementPortFlag, supernodeManagementPortKey, "m", 5645, "UDP management port")

	if err := viper.BindPFlags(supernodeCmd.PersistentFlags()); err != nil {
		log.Fatal(couldNotBindFlagsErrorMessage, rz.Err(err))
	}

	viper.AutomaticEnv()

	rootCmd.AddCommand(supernodeCmd)
}
