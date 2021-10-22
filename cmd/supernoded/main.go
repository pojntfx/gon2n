package main

import (
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	constants "github.com/pojntfx/gon2n/cmd"
	api "github.com/pojntfx/gon2n/pkg/api/proto/v1"
	"github.com/pojntfx/gon2n/pkg/services"
	"github.com/pojntfx/gon2n/pkg/workers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bloom42/libs/rz-go"
	"gitlab.com/bloom42/libs/rz-go/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	keyPrefix         = "supernoded."
	configFileDefault = ""
	configFileKey     = keyPrefix + "configFile"
	listenHostPortKey = keyPrefix + "listenHostPort"
)

var rootCmd = &cobra.Command{
	Use:   "supernoded",
	Short: "supernoded is the n2n supernode management daemon",
	Long: `supernoded is the n2n supernode management daemon.

Find more information at:
https://pojntfx.github.io/gon2n/`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		viper.SetEnvPrefix("supernoded")
		viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if !(viper.GetString(configFileKey) == configFileDefault) {
			viper.SetConfigFile(viper.GetString(configFileKey))

			if err := viper.ReadInConfig(); err != nil {
				return err
			}
		}

		listener, err := net.Listen("tcp", viper.GetString(listenHostPortKey))
		if err != nil {
			return err
		}

		server := grpc.NewServer()
		reflection.Register(server)

		supernodeService := services.SupernodeManager{
			SupernodesManaged: make(map[string]*workers.Supernode),
		}

		api.RegisterSupernodeManagerServer(server, &supernodeService)

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

			for _, supernode := range supernodeService.SupernodesManaged {
				if err := supernode.Stop(); err != nil {
					log.Fatal(msg, rz.Err(err))
				}
			}

			for _, supernode := range supernodeService.SupernodesManaged {
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
		configFileFlag string
		hostPortFlag   string
	)

	rootCmd.PersistentFlags().StringVarP(&configFileFlag, configFileKey, "f", configFileDefault, constants.ConfigurationFileDocs)
	rootCmd.PersistentFlags().StringVarP(&hostPortFlag, listenHostPortKey, "l", constants.SupernodeServerHostPortDefault, "TCP listen host:port.")

	if err := viper.BindPFlags(rootCmd.PersistentFlags()); err != nil {
		log.Fatal(constants.CouldNotBindFlagsErrorMessage, rz.Err(err))
	}

	viper.AutomaticEnv()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(constants.CouldNotStartRootCommandErrorMessage, rz.Err(err))
	}
}
