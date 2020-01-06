package cmd

import (
	"context"
	"fmt"
	"github.com/gosuri/uitable"
	gon2n "github.com/pojntfx/gon2n/pkg/proto/generated"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bloom42/libs/rz-go"
	"gitlab.com/bloom42/libs/rz-go/log"
	"google.golang.org/grpc"
)

var getEdgeCmd = &cobra.Command{
	Use:     "edge",
	Aliases: []string{"edges", "e"},
	Short:   "Get all edges",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := grpc.Dial(viper.GetString(edgeServerHostPortKey), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := gon2n.NewEdgeManagerClient(conn)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		response, err := client.List(ctx, &gon2n.EdgeManagerListArgs{})
		if err != nil {
			return err
		}

		table := uitable.New()
		table.AddRow(
			"ID",
			"ALLOW P2P",
			"ALLOW ROUTING",
			"COMMUNITY NAME",
			"DISABLE PATH MTU DISCOVERY",
			"DISABLE MULTICAST",
			"DYNAMIC IP MODE",
			"LOCAL PORT",
			"MANAGEMENT PORT",
			"REGISTER INTERVAL",
			"REGISTER TTL",
			"SUPERNODE HOST:PORT",
			"TYPE OF SERVICE",
			"ENCRYPTION METHOD",
			"DEVICE NAME",
			"ADDRESS MODE",
			"DEVICE IP",
			"DEVICE NETMASK",
			"DEVICE MAC ADDRESS",
			"MTU")

		for _, edge := range response.GetEdgesManaged() {
			table.AddRow(
				edge.GetId(),
				edge.GetAllowP2P(),
				edge.GetAllowRouting(),
				edge.GetCommunityName(),
				edge.GetDisablePMTUDiscovery(),
				edge.GetDisableMulticast(),
				edge.GetDynamicIPMode(),
				int(edge.GetLocalPort()),
				int(edge.GetManagementPort()),
				int(edge.GetRegisterInterval()),
				int(edge.GetRegisterTTL()),
				edge.GetSupernodeHostPort(),
				int(edge.GetTypeOfService()),
				int(edge.GetEncryptionMethod()),
				edge.GetDeviceName(),
				edge.GetAddressMode(),
				edge.GetDeviceIP(),
				edge.GetDeviceNetmask(),
				edge.GetDeviceMACAddress(),
				int(edge.GetMTU()))
		}

		fmt.Println(table)

		return nil
	},
}

func init() {
	var (
		edgeServerHostPortFlag string
	)

	getEdgeCmd.PersistentFlags().StringVarP(&edgeServerHostPortFlag, edgeServerHostPortKey, "s", "localhost:1235", "Host:port of the gon2n server to use.")

	if err := viper.BindPFlags(getEdgeCmd.PersistentFlags()); err != nil {
		log.Fatal(couldNotBindFlagsErrorMessage, rz.Err(err))
	}

	viper.AutomaticEnv()

	getCmd.AddCommand(getEdgeCmd)
}
