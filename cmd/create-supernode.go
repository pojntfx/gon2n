package cmd

import (
	"context"
	gon2n "github.com/pojntfx/gon2n/pkg/proto/generated/proto"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bloom42/libs/rz-go/v2"
	"gitlab.com/bloom42/libs/rz-go/v2/log"
	"google.golang.org/grpc"
	"time"
)

var supernodeCmd = &cobra.Command{
	Use:   "supernode",
	Short: "Create a supernode",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !(viper.GetString(supernodeConfigFileKey) == supernodeConfigFileDefault) {
			viper.SetConfigFile(viper.GetString(supernodeConfigFileKey))

			if err := viper.ReadInConfig(); err != nil {
				return err
			}
		}

		conn, err := grpc.Dial(viper.GetString(supernodeServerHostPortKey), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := gon2n.NewSupernodeManagerClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		log.Info("Starting supernode")

		response, err := client.Create(ctx, &gon2n.SupernodeManagerCreateArgs{
			ListenPort:     viper.GetInt64(supernodeListenPortKey),
			ManagementPort: viper.GetInt64(supernodeManagementPortKey),
		})
		if err != nil {
			return err
		}

		log.Info("Started supernode", rz.String("Id", response.GetId()))

		return nil
	},
}

func init() {
	var (
		supernodeServerHostPortFlag string
		supernodeConfigFileFlag     string
		supernodeListenPortFlag     int
		supernodeManagementPortFlag int
	)

	supernodeCmd.PersistentFlags().StringVarP(&supernodeServerHostPortFlag, supernodeServerHostPortKey, "s", "localhost:1235", "Host:port of the gon2n server to use.")
	supernodeCmd.PersistentFlags().StringVarP(&supernodeConfigFileFlag, supernodeConfigFileKey, "f", supernodeConfigFileDefault, "Configuration file to use.")
	supernodeCmd.PersistentFlags().IntVarP(&supernodeListenPortFlag, supernodeListenPortKey, "l", 1234, "UDP listen port.")
	supernodeCmd.PersistentFlags().IntVarP(&supernodeManagementPortFlag, supernodeManagementPortKey, "m", 5645, "UDP management port.")

	if err := viper.BindPFlags(supernodeCmd.PersistentFlags()); err != nil {
		log.Fatal(couldNotBindFlagsErrorMessage, rz.Err(err))
	}

	viper.AutomaticEnv()

	createCmd.AddCommand(supernodeCmd)
}
