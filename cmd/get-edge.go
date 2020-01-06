package cmd

import (
	"context"
	"fmt"

	"github.com/ghodss/yaml"
	"github.com/gosuri/uitable"
	gon2n "github.com/pojntfx/gon2n/pkg/proto/generated"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bloom42/libs/rz-go"
	"gitlab.com/bloom42/libs/rz-go/log"
	"google.golang.org/grpc"
)

var getEdgeCmd = &cobra.Command{
	Use:     "edge [id]",
	Aliases: []string{"edges", "e"},
	Short:   "Get one or all edge(s)",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := grpc.Dial(viper.GetString(edgeServerHostPortKey), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := gon2n.NewEdgeManagerClient(conn)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		if len(args) < 1 {
			response, err := client.List(ctx, &gon2n.EdgeManagerListArgs{})
			if err != nil {
				return err
			}

			table := uitable.New()
			table.AddRow(
				"ID",
				"COMMUNITY NAME",
				"LOCAL PORT",
				"SUPERNODE HOST:PORT",
				"ENCRYPTION METHOD",
				"DEVICE NAME")

			for _, edge := range response.GetEdgesManaged() {
				table.AddRow(
					edge.GetId(),
					edge.GetCommunityName(),
					edge.GetLocalPort(),
					edge.GetSupernodeHostPort(),
					edge.GetEncryptionMethod(),
					edge.GetDeviceName())
			}

			fmt.Println(table)

			return nil
		}

		response, err := client.Get(ctx, &gon2n.EdgeManagerGetArgs{
			Id: args[0],
		})
		if err != nil {
			return err
		}

		output, err := yaml.Marshal(&response)
		if err != nil {
			return err
		}

		fmt.Println(string(output))

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
