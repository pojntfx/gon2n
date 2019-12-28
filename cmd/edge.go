package cmd

import (
	"github.com/pojntfx/gon2n/pkg"
	"github.com/spf13/cobra"
)

var EdgeCmd = &cobra.Command{
	Use:   "edge",
	Short: "Start a n2n edge",
	Run: func(cmd *cobra.Command, args []string) {
		edge := pkg.Edge{
			DeviceName:    "edge1",
			NetworkName:   "dev",
			SecretKey:     "asdf",
			MyMacAddress:  "DE:AD:BE:EF:01:10",
			MyIPv4Addr:    "10.0.0.2",
			Supernode:     "192.168.178.35:1234",
			KeepOnRunning: true,
		}

		edge.Start()
	},
}

func init() {
	RootCmd.AddCommand(EdgeCmd)
}
