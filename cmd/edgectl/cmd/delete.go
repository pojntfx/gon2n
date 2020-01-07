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
	"sync"
)

var deleteCmd = &cobra.Command{
	Use:     "delete <id> [id...]",
	Aliases: []string{"d"},
	Short:   "Delete one or more edge(s)",
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := grpc.Dial(viper.GetString(serverHostPortKey), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := gon2n.NewEdgeManagerClient(conn)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var wg sync.WaitGroup

		for _, id := range args {
			wg.Add(1)

			go func(id string, wg *sync.WaitGroup) {
				response, err := client.Delete(ctx, &gon2n.EdgeManagerDeleteArgs{
					Id: id,
				})
				if err != nil {
					log.Error(err.Error())

					wg.Done()

					return
				}

				fmt.Printf("edge \"%s\" deleted\n", response.GetId())

				wg.Done()
			}(id, &wg)
		}

		wg.Wait()

		return nil
	},
}

func init() {
	var (
		serverHostPortFlag string
	)

	deleteCmd.PersistentFlags().StringVarP(&serverHostPortFlag, serverHostPortKey, "s", constants.EdgedServerHostPortDefault, "Host:port of the gon2n server to use.")

	if err := viper.BindPFlags(deleteCmd.PersistentFlags()); err != nil {
		log.Fatal(constants.CouldNotBindFlagsErrorMessage, rz.Err(err))
	}

	viper.AutomaticEnv()

	rootCmd.AddCommand(deleteCmd)
}
