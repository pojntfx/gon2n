package cmd

import (
	"context"
	"fmt"
	gon2n "github.com/pojntfx/gon2n/pkg/proto/generated"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bloom42/libs/rz-go"
	"gitlab.com/bloom42/libs/rz-go/log"
	"google.golang.org/grpc"
	"sync"
)

var deleteSupernodeCmd = &cobra.Command{
	Use:     "supernode <id> [id...]",
	Aliases: []string{"supernodes", "s"},
	Short:   "Delete one or more supernode(s)",
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := grpc.Dial(viper.GetString(supernodeServerHostPortKey), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := gon2n.NewSupernodeManagerClient(conn)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var wg sync.WaitGroup

		for _, id := range args {
			wg.Add(1)

			go func(id string, wg *sync.WaitGroup) {
				response, err := client.Delete(ctx, &gon2n.SupernodeManagerDeleteArgs{
					Id: id,
				})
				if err != nil {
					log.Error(err.Error())

					wg.Done()

					return
				}

				fmt.Printf("supernode \"%s\" deleted\n", response.GetId())

				wg.Done()
			}(id, &wg)
		}

		wg.Wait()

		return nil
	},
}

func init() {
	var (
		supernodeServerHostPortFlag string
	)

	deleteSupernodeCmd.PersistentFlags().StringVarP(&supernodeServerHostPortFlag, supernodeServerHostPortKey, "s", "localhost:1235", "Host:port of the gon2n server to use.")

	if err := viper.BindPFlags(deleteSupernodeCmd.PersistentFlags()); err != nil {
		log.Fatal(couldNotBindFlagsErrorMessage, rz.Err(err))
	}

	viper.AutomaticEnv()

	deleteCmd.AddCommand(deleteSupernodeCmd)
}
