package cmd

import (
	"github.com/pojntfx/gon2n/pkg"
	"github.com/spf13/cobra"
)

var SupernodeCmd = &cobra.Command{
	Use:   "supernode",
	Short: "Start a supernode",
	RunE: func(cmd *cobra.Command, args []string) error {
		supernode := pkg.Supernode{
			ListenPort:     listenPort,
			ManagementPort: managementPort,
		}

		return supernode.Start()
	},
}

var (
	listenPort     int
	managementPort int
)

func init() {
	SupernodeCmd.PersistentFlags().IntVarP(&listenPort, "port-listen", "l", 1234, "Main UDP listen port")
	SupernodeCmd.PersistentFlags().IntVarP(&managementPort, "port-management", "m", 5645, "Main UDP management port")

	RootCmd.AddCommand(SupernodeCmd)
}
