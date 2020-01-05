package cmd

import (
	gon2n "github.com/pojntfx/gon2n/pkg/proto/generated/proto"
	"github.com/pojntfx/gon2n/pkg/svc"
	"github.com/pojntfx/gon2n/pkg/workers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bloom42/libs/rz-go/v2"
	"gitlab.com/bloom42/libs/rz-go/v2/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var serverCmd = &cobra.Command{
	Use:     "server",
	Aliases: []string{"s"},
	Short:   "Start a gon2n server",
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

		service := svc.SupernodeManager{
			SupernodesManaged: make(map[string]*workers.Supernode),
		}

		gon2n.RegisterSupernodeManagerServer(server, &service)

		interrupt := make(chan os.Signal, 2)
		signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-interrupt

			// Allow manually killing the process
			go func() {
				<-interrupt

				os.Exit(1)
			}()

			log.Info("Gracefully stopping server (this might take a few seconds)")

			msg := "Could not stop supernode"

			for _, supernode := range service.SupernodesManaged {
				if err := supernode.Stop(); err != nil {
					log.Fatal(msg, rz.Err(err))
				}
			}

			for _, supernode := range service.SupernodesManaged {
				if err := supernode.Wait(); err != nil {
					log.Fatal(msg, rz.Err(err))
				}
			}

			server.GracefulStop()
		}()

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
	serverCmd.PersistentFlags().StringVarP(&serverHostPortFlag, serverListenHostPortKey, "l", "localhost:1235", "TCP listen host:port.")

	if err := viper.BindPFlags(serverCmd.PersistentFlags()); err != nil {
		log.Fatal(couldNotBindFlagsErrorMessage, rz.Err(err))
	}

	viper.AutomaticEnv()

	rootCmd.AddCommand(serverCmd)
}
