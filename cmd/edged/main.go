package main

import (
	constants "github.com/pojntfx/gon2n/cmd"
	gon2n "github.com/pojntfx/gon2n/pkg/proto/generated"
	"github.com/pojntfx/gon2n/pkg/svc"
	"github.com/pojntfx/gon2n/pkg/workers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bloom42/libs/rz-go"
	"gitlab.com/bloom42/libs/rz-go/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

const (
	keyPrefix         = "edged."
	configFileDefault = ""
	configFileKey     = keyPrefix + "configFile"
	listenHostPortKey = keyPrefix + "listenHostPort"
)

var rootCmd = &cobra.Command{
	Use:   "edged",
	Short: "edged is the n2n edge management daemon",
	Long: `edged is the n2n edge management daemon.

Find more information at:
https://pojntfx.github.io/gon2n/`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		viper.SetEnvPrefix("edged")
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

		edgeService := svc.EdgeManager{
			EdgesManaged: make(map[string]*workers.Edge),
		}

		gon2n.RegisterEdgeManagerServer(server, &edgeService)

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

			msg := "Could not stop edge"

			for _, edge := range edgeService.EdgesManaged {
				if err := edge.Stop(); err != nil {
					log.Fatal(msg, rz.Err(err))
				}
			}

			for _, edge := range edgeService.EdgesManaged {
				if err := edge.Wait(); err != nil {
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
	rootCmd.PersistentFlags().StringVarP(&hostPortFlag, listenHostPortKey, "l", constants.EdgedServerHostPortDefault, "TCP listen host:port.")

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
