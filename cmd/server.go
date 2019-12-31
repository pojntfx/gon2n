package cmd

import (
	"net"

	gon2n "github.com/pojntfx/gon2n/pkg/proto/generated/proto"
	"github.com/pojntfx/gon2n/pkg/svc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bloom42/libs/rz-go/v2"
	"gitlab.com/bloom42/libs/rz-go/v2/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a gon2n server",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !(viper.GetString(serverConfigFileKey) == serverConfigFileDefault) {
			viper.SetConfigFile(viper.GetString(serverConfigFileKey))

			if err := viper.ReadInConfig(); err != nil {
				return err
			}
		}

		listener, err := net.Listen("tcp", viper.GetString(serverListenHostPortKey))
		if err != nil {
			return err
		}

		server := grpc.NewServer()
		reflection.Register(server)

		gon2n.RegisterSupernodeManagerServer(server, &svc.SupernodeManager{})

		log.Info("Starting server")

		return server.Serve(listener)
	},
}

func init() {
	var (
		serverConfigFileFlag string
		serverHostPortFlag   string
	)

	serverCmd.PersistentFlags().StringVarP(&serverConfigFileFlag, serverConfigFileKey, "f", serverConfigFileDefault, "Configuration file to use.")
	serverCmd.PersistentFlags().StringVarP(&serverHostPortFlag, serverListenHostPortKey, "l", "127.0.0.1:1235", "TCP listen host:port.")

	if err := viper.BindPFlags(serverCmd.PersistentFlags()); err != nil {
		log.Fatal(couldNotBindFlagsErrorMessage, rz.Err(err))
	}

	viper.AutomaticEnv()

	rootCmd.AddCommand(serverCmd)
}
