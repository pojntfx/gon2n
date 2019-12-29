package cmd

import (
	"github.com/pojntfx/gon2n/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/z0mbie42/rz-go/v2"
	"gitlab.com/z0mbie42/rz-go/v2/log"
)

var edgeCmd = &cobra.Command{
	Use:   "edge",
	Short: "Start an edge",
	RunE: func(cmd *cobra.Command, args []string) error {
		edge := pkg.Edge{
			AllowP2P:             viper.GetBool(edgeAllowP2PKey),
			AllowRouting:         viper.GetBool(edgeAllowRoutingKey),
			CommunityName:        viper.GetString(edgeCommunityNameKey),
			DisablePMTUDiscovery: viper.GetBool(edgeDisablePMTUDiscoveryKey),
			DisableMulticast:     viper.GetBool(edgeDisableMulticastKey),
			DynamicIPMode:        viper.GetBool(edgeDynamicIPModeKey),
			EncryptionKey:        viper.GetString(edgeEncryptionKeyKey),
			LocalPort:            viper.GetInt(edgeLocalPortKey),
			ManagementPort:       viper.GetInt(edgeManagementPortKey),
			RegisterInterval:     viper.GetInt(edgeRegisterIntervalKey),
			RegisterTTL:          viper.GetInt(edgeRegisterTTLKey),
			SupernodeHostPort:    viper.GetString(edgeSupernodeHostPortKey),
			TypeOfService:        viper.GetInt(edgeTypeOfServiceKey),
			EncryptionMethod:     viper.GetInt(edgeEncryptionMethodKey),
			DeviceName:           viper.GetString(edgeDeviceNameKey),
			AddressMode:          viper.GetString(edgeAddressModeKey),
			DeviceIP:             viper.GetString(edgeDeviceIPKey),
			DeviceNetmask:        viper.GetString(edgeDeviceNetmaskKey),
			DeviceMACAddress:     viper.GetString(edgeDeviceMACAddressKey),
			MTU:                  viper.GetInt(edgeMTUKey),
		}

		return edge.Start()
	},
}

func init() {
	var (
		edgeAllowP2PFlag             bool
		edgeAllowRoutingFlag         bool
		edgeCommunityNameFlag        string
		edgeDisablePMTUDiscoveryFlag bool
		edgeDisableMulticastFlag     bool
		edgeDynamicIPModeFlag        bool
		edgeEncryptionKeyFlag        string
		edgeLocalPortFlag            int
		edgeManagementPortFlag       int
		edgeRegisterIntervalFlag     int
		edgeRegisterTTLFlag          int
		edgeSupernodeHostPortFlag    string
		edgeTypeOfServiceFlag        int
		edgeEncryptionMethodFlag     int
		edgeDeviceNameFlag           string
		edgeAddressModeFlag          string
		edgeDeviceIPFlag             string
		edgeDeviceNetmaskFlag        string
		edgeDeviceMACAddressFlag     string
		edgeMTUFlag                  int
	)

	edgeCmd.PersistentFlags().BoolVarP(&edgeAllowP2PFlag, edgeAllowP2PKey, "p", true, "Whether to allow peer-to-peer connections. If false, all traffic will be routed through the supernode.")
	edgeCmd.PersistentFlags().BoolVarP(&edgeAllowRoutingFlag, edgeAllowRoutingKey, "r", true, "Whether to allow the node to route traffic to other nodes.")
	edgeCmd.PersistentFlags().StringVarP(&edgeCommunityNameFlag, edgeCommunityNameKey, "c", "mynetwork", "The name of the n2n community to join.")
	edgeCmd.PersistentFlags().BoolVarP(&edgeDisablePMTUDiscoveryFlag, edgeDisablePMTUDiscoveryKey, "d", false, "Whether to allow peer-to-peer connections. If false, all traffic will be routed through the supernode.")
	edgeCmd.PersistentFlags().BoolVarP(&edgeDisableMulticastFlag, edgeDisableMulticastKey, "m", false, "Whether to disable multicast.")
	edgeCmd.PersistentFlags().BoolVarP(&edgeDynamicIPModeFlag, edgeDynamicIPModeKey, "y", false, "Whether the IP address is set dynamically (see --address-mode). If the edge is running the network's DHCP server, this must be false.")
	edgeCmd.PersistentFlags().StringVarP(&edgeEncryptionKeyFlag, edgeEncryptionKeyKey, "k", "mysecretkey", "The key to use for encryption.")
	edgeCmd.PersistentFlags().IntVarP(&edgeLocalPortFlag, edgeLocalPortKey, "l", 0, "The local port to use. 0 uses any available port.")
	edgeCmd.PersistentFlags().IntVarP(&edgeManagementPortFlag, edgeManagementPortKey, "a", 5644, "UDP management port. 5644 is the n2n default.")
	edgeCmd.PersistentFlags().IntVarP(&edgeRegisterIntervalFlag, edgeRegisterIntervalKey, "n", 1, "Interval in seconds for both UDP NAT hole punching and registration of the edge on the supernode.")
	edgeCmd.PersistentFlags().IntVarP(&edgeRegisterTTLFlag, edgeRegisterTTLKey, "t", 1, "Interval in seconds for UDP NAT hole punching through the supernode.")
	edgeCmd.PersistentFlags().StringVarP(&edgeSupernodeHostPortFlag, edgeSupernodeHostPortKey, "s", "localhost:1234", "Host:port of the supernode to connect to.")
	edgeCmd.PersistentFlags().IntVarP(&edgeTypeOfServiceFlag, edgeTypeOfServiceKey, "o", 16, "Type of service to use.")
	edgeCmd.PersistentFlags().IntVarP(&edgeEncryptionMethodFlag, edgeEncryptionMethodKey, "e", 2, "Method of encryption to use. 1 is no encryption, 2 is Twofish encryption, 3 is AES-CBC encryption. Twofish encryption is the n2n default, but only due to legacy compatibility reasons; AES-CBC is up to ten times faster.")
	edgeCmd.PersistentFlags().StringVarP(&edgeDeviceNameFlag, edgeDeviceNameKey, "v", "edge0", "Name of the TUN/TAP device to create.")
	edgeCmd.PersistentFlags().StringVarP(&edgeAddressModeFlag, edgeAddressModeKey, "z", "static", "Mode of IP address assignment. \"static\" is a static assignment, \"dhcp\" uses the DHCP server at --device-ip (see --dynamic-ip-node). If the edge is running the network's DHCP server, this must be \"static\".")
	edgeCmd.PersistentFlags().StringVarP(&edgeDeviceIPFlag, edgeDeviceIPKey, "i", "10.0.0.1", "IP address to set. Set to \"0.0.0.0\" if you are using \"dhcp\" as --address-mode. If the edge is running the network's DHCP server this must be set explicitly; i.e. to \"192.168.1.0\" if the DHCP server should give out addresses from \"192.168.1.10\" to \"192.168.1.100\".")
	edgeCmd.PersistentFlags().StringVarP(&edgeDeviceNetmaskFlag, edgeDeviceNetmaskKey, "q", "255.255.255.0", "The netmask to use.")
	edgeCmd.PersistentFlags().StringVarP(&edgeDeviceMACAddressFlag, edgeDeviceMACAddressKey, "x", "DE:AD:BE:EF:01:10", "The MAC address to use. Must be unique per edge.")
	edgeCmd.PersistentFlags().IntVarP(&edgeMTUFlag, edgeMTUKey, "u", 1290, "The MTU to use.")

	if err := viper.BindPFlags(edgeCmd.PersistentFlags()); err != nil {
		log.Fatal(couldNotBindFlagsErrorMessage, rz.Err(err))
	}

	viper.AutomaticEnv()

	rootCmd.AddCommand(edgeCmd)
}
