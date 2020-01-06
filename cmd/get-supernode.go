package cmd

import (
	"context"
	"fmt"
	"github.com/gosuri/uitable"
	gon2n "github.com/pojntfx/gon2n/pkg/proto/generated/proto"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bloom42/libs/rz-go/v2"
	"gitlab.com/bloom42/libs/rz-go/v2/log"
	"google.golang.org/grpc"
)

var getSupernodeCmd = &cobra.Command{
	Use:     "supernode",
	Aliases: []string{"supernodes", "s"},
	Short:   "Get all supernodes",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := grpc.Dial(viper.GetString(supernodeServerHostPortKey), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := gon2n.NewSupernodeManagerClient(conn)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		response, err := client.List(ctx, &gon2n.SupernodeManagerListArgs{})
		if err != nil {
			return err
		}

		table := uitable.New()
		table.AddRow("ID", "LISTEN PORT", "MANAGEMENT PORT")

		for _, supernode := range response.GetSupernodesManaged() {
			table.AddRow(supernode.GetId(), supernode.GetListenPort(), supernode.GetManagementPort())
		}

		fmt.Println(table)

		return nil
	},
}

func init() {
	var (
		supernodeServerHostPortFlag string
	)

	getSupernodeCmd.PersistentFlags().StringVarP(&supernodeServerHostPortFlag, supernodeServerHostPortKey, "s", "localhost:1235", "Host:port of the gon2n server to use.")

	if err := viper.BindPFlags(getSupernodeCmd.PersistentFlags()); err != nil {
		log.Fatal(couldNotBindFlagsErrorMessage, rz.Err(err))
	}

	viper.AutomaticEnv()

	getCmd.AddCommand(getSupernodeCmd)
}
