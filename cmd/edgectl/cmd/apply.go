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
)

var applyCmd = &cobra.Command{
	Use:     "apply",
	Aliases: []string{"a"},
	Short:   "Apply an edge",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !(viper.GetString(configFileKey) == configFileDefault) {
			viper.SetConfigFile(viper.GetString(configFileKey))

			if err := viper.ReadInConfig(); err != nil {
				return err
			}
		}

		conn, err := grpc.Dial(viper.GetString(serverHostPortKey), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := gon2n.NewEdgeManagerClient(conn)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		response, err := client.Create(ctx, &gon2n.Edge{
			AllowP2P:             viper.GetBool(allowP2PKey),
			AllowRouting:         viper.GetBool(allowRoutingKey),
			CommunityName:        viper.GetString(communityNameKey),
			DisablePMTUDiscovery: viper.GetBool(disablePMTUDiscoveryKey),
			DisableMulticast:     viper.GetBool(disableMulticastKey),
			DynamicIPMode:        viper.GetBool(dynamicIPModeKey),
			EncryptionKey:        viper.GetString(encryptionKeyKey),
			LocalPort:            int64(viper.GetInt(localPortKey)),
			ManagementPort:       int64(viper.GetInt(managementPortKey)),
			RegisterInterval:     int64(viper.GetInt(registerIntervalKey)),
			RegisterTTL:          int64(viper.GetInt(registerTTLKey)),
			SupernodeHostPort:    viper.GetString(supernodeHostPortKey),
			TypeOfService:        int64(viper.GetInt(typeOfServiceKey)),
			EncryptionMethod:     int64(viper.GetInt(encryptionMethodKey)),
			DeviceName:           viper.GetString(deviceNameKey),
			AddressMode:          viper.GetString(addressModeKey),
			DeviceIP:             viper.GetString(deviceIPKey),
			DeviceNetmask:        viper.GetString(deviceNetmaskKey),
			DeviceMACAddress:     viper.GetString(deviceMACAddressKey),
			MTU:                  int64(viper.GetInt(mtuKey)),
		})
		if err != nil {
			return err
		}

		fmt.Printf("edge \"%s\" created\n", response.GetId())

		return nil
	},
}

func init() {
	var (
		serverHostPortFlag       string
		configFileFlag           string
		allowP2PFlag             bool
		allowRoutingFlag         bool
		communityNameFlag        string
		disablePMTUDiscoveryFlag bool
		disableMulticastFlag     bool
		dynamicIPModeFlag        bool
		encryptionKeyFlag        string
		localPortFlag            int
		managementPortFlag       int
		registerIntervalFlag     int
		registerTTLFlag          int
		supernodeHostPortFlag    string
		typeOfServiceFlag        int
		encryptionMethodFlag     int
		deviceNameFlag           string
		addressModeFlag          string
		deviceIPFlag             string
		deviceNetmaskFlag        string
		deviceMACAddressFlag     string
		mtuFlag                  int
	)

	applyCmd.PersistentFlags().StringVarP(&serverHostPortFlag, serverHostPortKey, "s", constants.EdgedServerHostPortDefault, "Host:port of the gon2n server to use.")
	applyCmd.PersistentFlags().StringVarP(&configFileFlag, configFileKey, "f", configFileDefault, "Configuration file to use.")
	applyCmd.PersistentFlags().BoolVarP(&allowP2PFlag, allowP2PKey, "p", true, "Whether to allow peer-to-peer connections. If false, all traffic will be routed through the supernode.")
	applyCmd.PersistentFlags().BoolVarP(&allowRoutingFlag, allowRoutingKey, "r", true, "Whether to allow the node to route traffic to other nodes.")
	applyCmd.PersistentFlags().StringVarP(&communityNameFlag, communityNameKey, "c", "mynetwork", "The name of the n2n community to join.")
	applyCmd.PersistentFlags().BoolVarP(&disablePMTUDiscoveryFlag, disablePMTUDiscoveryKey, "d", false, "Whether to disable path MTU discovery.")
	applyCmd.PersistentFlags().BoolVarP(&disableMulticastFlag, disableMulticastKey, "m", false, "Whether to disable multicast.")
	applyCmd.PersistentFlags().BoolVarP(&dynamicIPModeFlag, dynamicIPModeKey, "y", false, "Whether the IP address is set dynamically (see --address-mode). If the edge is running the network's DHCP server, this must be false.")
	applyCmd.PersistentFlags().StringVarP(&encryptionKeyFlag, encryptionKeyKey, "k", "mysecretkey", "The key to use for encryption.")
	applyCmd.PersistentFlags().IntVarP(&localPortFlag, localPortKey, "l", 0, "The local port to use. 0 uses any available port.")
	applyCmd.PersistentFlags().IntVarP(&managementPortFlag, managementPortKey, "a", 5644, "UDP management port. 5644 is the n2n default.")
	applyCmd.PersistentFlags().IntVarP(&registerIntervalFlag, registerIntervalKey, "n", 1, "Interval in seconds for both UDP NAT hole punching and registration of the edge on the supernode.")
	applyCmd.PersistentFlags().IntVarP(&registerTTLFlag, registerTTLKey, "t", 1, "Interval in seconds for UDP NAT hole punching through the supernode.")
	applyCmd.PersistentFlags().StringVarP(&supernodeHostPortFlag, supernodeHostPortKey, "w", constants.SupernodeHostPortDefault, "Host:port of the supernode to connect to.")
	applyCmd.PersistentFlags().IntVarP(&typeOfServiceFlag, typeOfServiceKey, "o", 16, "Type of service to use.")
	applyCmd.PersistentFlags().IntVarP(&encryptionMethodFlag, encryptionMethodKey, "e", 2, "Method of encryption to use. 1 is no encryption, 2 is Twofish encryption, 3 is AES-CBC encryption. Twofish encryption is the n2n default, but only due to legacy compatibility reasons; AES-CBC is up to ten times faster.")
	applyCmd.PersistentFlags().StringVarP(&deviceNameFlag, deviceNameKey, "v", "edge0", "Name of the TUN/TAP device to create.")
	applyCmd.PersistentFlags().StringVarP(&addressModeFlag, addressModeKey, "z", "static", "Mode of IP address assignment. \"static\" is a static assignment, \"dhcp\" uses the DHCP server at --device-ip (see --dynamic-ip-node). If the edge is running the network's DHCP server, this must be \"static\".")
	applyCmd.PersistentFlags().StringVarP(&deviceIPFlag, deviceIPKey, "i", "192.168.1.1", "IP address to set. Set to \"0.0.0.0\" if you are using \"dhcp\" as --address-mode. If the edge is running the network's DHCP server this must be set explicitly; i.e. to \"192.168.1.0\" if the DHCP server should give out addresses from \"192.168.1.10\" to \"192.168.1.100\".")
	applyCmd.PersistentFlags().StringVarP(&deviceNetmaskFlag, deviceNetmaskKey, "q", "255.255.255.0", "The netmask to use.")
	applyCmd.PersistentFlags().StringVarP(&deviceMACAddressFlag, deviceMACAddressKey, "x", "DE:AD:BE:EF:01:10", "The MAC address to use. Must be unique per edge.")
	applyCmd.PersistentFlags().IntVarP(&mtuFlag, mtuKey, "u", 1500, "The MTU to use.")

	if err := viper.BindPFlags(applyCmd.PersistentFlags()); err != nil {
		log.Fatal(constants.CouldNotBindFlagsErrorMessage, rz.Err(err))
	}

	viper.AutomaticEnv()

	rootCmd.AddCommand(applyCmd)
}
