package cmd

import (
	"context"
	"fmt"
	constants "github.com/pojntfx/gon2n/cmd"
	gon2n "github.com/pojntfx/gon2n/pkg/proto/generated"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bloom42/libs/rz-go"
	"gitlab.com/bloom42/libs/rz-go/log"
	"google.golang.org/grpc"
)

var applyCmd = &cobra.Command{
	Use:     "apply",
	Aliases: []string{"a"},
	Short:   "Apply a supernode",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !(viper.GetString(configFileKey) == configFileDefault) {
			viper.SetConfigFile(viper.GetString(configFileKey))

			if err := viper.ReadInConfig(); err != nil {
				return err
			}
		}

		conn, err := grpc.Dial(viper.GetString(serverHostPortKey), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := gon2n.NewSupernodeManagerClient(conn)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		response, err := client.Create(ctx, &gon2n.Supernode{
			ListenPort:     viper.GetInt64(listenPortKey),
			ManagementPort: viper.GetInt64(managementPortKey),
		})
		if err != nil {
			return err
		}

		fmt.Printf("supernode \"%s\" created\n", response.GetId())

		return nil
	},
}

func init() {
	var (
		serverHostPortFlag string
		configFileFlag     string
		listenPortFlag     int
		managementPortFlag int
	)

	applyCmd.PersistentFlags().StringVarP(&serverHostPortFlag, serverHostPortKey, "s", constants.SupernodeServerHostPortDefault, constants.HostPortDocs)
	applyCmd.PersistentFlags().StringVarP(&configFileFlag, configFileKey, "f", configFileDefault, constants.ConfigurationFileDocs)
	applyCmd.PersistentFlags().IntVarP(&listenPortFlag, listenPortKey, "l", 1234, "UDP listen port.")
	applyCmd.PersistentFlags().IntVarP(&managementPortFlag, managementPortKey, "m", 5645, "UDP management port.")

	if err := viper.BindPFlags(applyCmd.PersistentFlags()); err != nil {
		log.Fatal(constants.CouldNotBindFlagsErrorMessage, rz.Err(err))
	}

	viper.AutomaticEnv()

	rootCmd.AddCommand(applyCmd)
}
