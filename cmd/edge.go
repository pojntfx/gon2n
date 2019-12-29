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
			DeviceName:           "edge0",
			AddressMode:          "static",
			DeviceIP:             "10.0.0.1",
			DeviceNetmask:        "255.255.255.0",
			DeviceMACAddress:     "DE:AD:BE:EF:01:10",
			MTU:                  1290,
		}

		return edge.Start()
	},
}

func init() {
	RootCmd.AddCommand(EdgeCmd)
}
