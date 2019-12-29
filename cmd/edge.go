package cmd

import (
	"github.com/pojntfx/gon2n/pkg"
	"github.com/spf13/cobra"
)

var EdgeCmd = &cobra.Command{
	Use:   "edge",
	Short: "Start an edge",
	RunE: func(cmd *cobra.Command, args []string) error {
		edge := pkg.Edge{
			AllowP2P:             true,
			AllowRouting:         true,
			CommunityName:        "mynetwork",
			DisablePMTUDiscovery: false,
			DropMulticast:        false,
			DynamicIPMode:        false,
			EncryptionKey:        "mysecret",
			LocalPort:            0,
			ManagementPort:       5644,
			RegisterInterval:     1,
			RegisterTTL:          1,
			SupernodeHostPort:    "localhost:1234",
			TypeOfService:        16,
			EncryptionMethod:     2,
		}

		return edge.Start()
	},
}

func init() {
	RootCmd.AddCommand(EdgeCmd)
}
