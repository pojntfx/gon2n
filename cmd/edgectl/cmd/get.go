package cmd

import (
	"context"
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/gosuri/uitable"
	constants "github.com/pojntfx/gon2n/cmd"
	gon2n "github.com/pojntfx/gon2n/pkg/proto/generated"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bloom42/libs/rz-go"
	"gitlab.com/bloom42/libs/rz-go/log"
	"google.golang.org/grpc"
)

var getCmd = &cobra.Command{
	Use:     "get [id]",
	Aliases: []string{"g"},
	Short:   "Get one or all edge(s)",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := grpc.Dial(viper.GetString(serverHostPortKey), grpc.WithInsecure(), grpc.WithBlock())
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

		response, err := client.Get(ctx, &gon2n.EdgeManagedId{
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
		serverHostPortFlag string
	)

	getCmd.PersistentFlags().StringVarP(&serverHostPortFlag, serverHostPortKey, "s", constants.EdgedServerHostPortDefault, constants.HostPortDocs)

	if err := viper.BindPFlags(getCmd.PersistentFlags()); err != nil {
		log.Fatal(constants.CouldNotBindFlagsErrorMessage, rz.Err(err))
	}

	viper.AutomaticEnv()

	rootCmd.AddCommand(getCmd)
}
