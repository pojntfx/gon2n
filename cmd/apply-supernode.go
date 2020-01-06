package cmd

import (
	"context"
	"fmt"
	gon2n "github.com/pojntfx/gon2n/pkg/proto/generated/proto"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bloom42/libs/rz-go"
	"gitlab.com/bloom42/libs/rz-go/log"
	"google.golang.org/grpc"
)

var applySupernodeCmd = &cobra.Command{
	Use:     "supernode",
	Aliases: []string{"supernodes", "s"},
	Short:   "Apply a supernode",
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

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		response, err := client.Create(ctx, &gon2n.SupernodeManagerCreateArgs{
			ListenPort:     viper.GetInt64(supernodeListenPortKey),
			ManagementPort: viper.GetInt64(supernodeManagementPortKey),
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
		supernodeServerHostPortFlag string
		supernodeConfigFileFlag     string
		supernodeListenPortFlag     int
		supernodeManagementPortFlag int
	)

	applySupernodeCmd.PersistentFlags().StringVarP(&supernodeServerHostPortFlag, supernodeServerHostPortKey, "s", "localhost:1235", "Host:port of the gon2n server to use.")
	applySupernodeCmd.PersistentFlags().StringVarP(&supernodeConfigFileFlag, supernodeConfigFileKey, "f", supernodeConfigFileDefault, "Configuration file to use.")
	applySupernodeCmd.PersistentFlags().IntVarP(&supernodeListenPortFlag, supernodeListenPortKey, "l", 1234, "UDP listen port.")
	applySupernodeCmd.PersistentFlags().IntVarP(&supernodeManagementPortFlag, supernodeManagementPortKey, "m", 5645, "UDP management port.")

	if err := viper.BindPFlags(applySupernodeCmd.PersistentFlags()); err != nil {
		log.Fatal(couldNotBindFlagsErrorMessage, rz.Err(err))
	}

	viper.AutomaticEnv()

	applyCmd.AddCommand(applySupernodeCmd)
}
