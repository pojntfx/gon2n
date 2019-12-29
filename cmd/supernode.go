package cmd

import (
	"github.com/pojntfx/gon2n/pkg"
	"github.com/spf13/cobra"
)

// SupernodeCmd starts a supernode.
var SupernodeCmd = &cobra.Command{
	Use:   "supernode",
	Short: "Start a supernode",
	RunE: func(cmd *cobra.Command, args []string) error {
		supernode := pkg.Supernode{
			ListenPort:     supernodeListenPort,
			ManagementPort: supernodeManagementPort,
		}

		return supernode.Start()
	},
}

var (
	supernodeListenPort     int
	supernodeManagementPort int
)

func init() {
	SupernodeCmd.PersistentFlags().IntVarP(&supernodeListenPort, "listen-port", "l", 1234, "UDP listen port")
	SupernodeCmd.PersistentFlags().IntVarP(&supernodeManagementPort, "management-port", "m", 5645, "UDP management port")

	RootCmd.AddCommand(SupernodeCmd)
}
