package cmd

import (
	"github.com/pojntfx/gon2n/pkg"
	"github.com/spf13/cobra"
)

var edgeCmd = &cobra.Command{
	Use:   "edge",
	Short: "Start an edge",
	RunE: func(cmd *cobra.Command, args []string) error {
		edge := pkg.Edge{
			AllowP2P:             edgeAllowP2P,
			AllowRouting:         edgeAllowRouting,
			CommunityName:        edgeCommunityName,
			DisablePMTUDiscovery: edgeDisablePMTUDiscovery,
			DisableMulticast:     edgeDisableMulticast,
			DynamicIPMode:        edgeDynamicIPMode,
			EncryptionKey:        edgeEncryptionKey,
			LocalPort:            edgeLocalPort,
			ManagementPort:       edgeManagementPort,
			RegisterInterval:     edgeRegisterInterval,
			RegisterTTL:          edgeRegisterTTL,
			SupernodeHostPort:    edgeSupernodeHostPort,
			TypeOfService:        edgeTypeOfService,
			EncryptionMethod:     edgeEncryptionMethod,
			DeviceName:           edgeDeviceName,
			AddressMode:          edgeAddressMode,
			DeviceIP:             edgeDeviceIP,
			DeviceNetmask:        edgeDeviceNetmask,
			DeviceMACAddress:     edgeDeviceMACAddress,
			MTU:                  edgeMTU,
		}

		return edge.Start()
	},
}

var (
	edgeAllowP2P             bool
	edgeAllowRouting         bool
	edgeCommunityName        string
	edgeDisablePMTUDiscovery bool
	edgeDisableMulticast     bool
	edgeDynamicIPMode        bool
	edgeEncryptionKey        string
	edgeLocalPort            int
	edgeManagementPort       int
	edgeRegisterInterval     int
	edgeRegisterTTL          int
	edgeSupernodeHostPort    string
	edgeTypeOfService        int
	edgeEncryptionMethod     int
	edgeDeviceName           string
	edgeAddressMode          string
	edgeDeviceIP             string
	edgeDeviceNetmask        string
	edgeDeviceMACAddress     string
	edgeMTU                  int
)

func init() {
	edgeCmd.PersistentFlags().BoolVarP(&edgeAllowP2P, "allow-p2p", "p", true, "Whether to allow peer-to-peer connections. If false, all traffic will be routed through the supernode.")
	edgeCmd.PersistentFlags().BoolVarP(&edgeAllowRouting, "allow-routing", "r", true, "Whether to allow the node to route traffic to other nodes.")
	edgeCmd.PersistentFlags().StringVarP(&edgeCommunityName, "community-name", "c", "mynetwork", "The name of the n2n community to join.")
	edgeCmd.PersistentFlags().BoolVarP(&edgeDisablePMTUDiscovery, "disable-pmtu-discovery", "d", false, "Whether to allow peer-to-peer connections. If false, all traffic will be routed through the supernode.")
	edgeCmd.PersistentFlags().BoolVarP(&edgeDisableMulticast, "disable-multicast", "m", false, "Whether to disable multicast.")
	edgeCmd.PersistentFlags().BoolVarP(&edgeDynamicIPMode, "dynamic-ip-mode", "y", false, "Whether the IP address is set dynamically (see --address-mode). If the edge is running the network's DHCP server, this must be false.")
	edgeCmd.PersistentFlags().StringVarP(&edgeEncryptionKey, "encryption-key", "k", "mysecretkey", "The key to use for encryption.")
	edgeCmd.PersistentFlags().IntVarP(&edgeLocalPort, "local-port", "l", 0, "The local port to use. 0 uses any available port.")
	edgeCmd.PersistentFlags().IntVarP(&edgeManagementPort, "management-port", "a", 5644, "UDP management port. 5644 is the n2n default.")
	edgeCmd.PersistentFlags().IntVarP(&edgeRegisterInterval, "register-interval", "n", 1, "Interval in seconds for both UDP NAT hole punching and registration of the edge on the supernode.")
	edgeCmd.PersistentFlags().IntVarP(&edgeRegisterTTL, "register-ttl", "t", 1, "Interval in seconds for UDP NAT hole punching through the supernode.")
	edgeCmd.PersistentFlags().StringVarP(&edgeSupernodeHostPort, "supernode-host-port", "s", "localhost:1234", "Host:port of the supernode to connect to.")
	edgeCmd.PersistentFlags().IntVarP(&edgeTypeOfService, "type-of-service", "o", 16, "Type of service to use.")
	edgeCmd.PersistentFlags().IntVarP(&edgeEncryptionMethod, "encryption-method", "e", 2, "Method of encryption to use. 1 is no encryption, 2 is Twofish encryption, 3 is AES-CBC encryption. Twofish encryption is the n2n default, but only due to legacy compatibility reasons; AES-CBC is up to ten times faster.")
	edgeCmd.PersistentFlags().StringVarP(&edgeDeviceName, "device-name", "v", "edge0", "Name of the TUN/TAP device to create.")
	edgeCmd.PersistentFlags().StringVarP(&edgeAddressMode, "address-mode", "z", "static", "Mode of IP address assignment. \"static\" is a static assignment, \"dhcp\" uses the DHCP server at --device-ip (see --dynamic-ip-node). If the edge is running the network's DHCP server, this must be \"static\".")
	edgeCmd.PersistentFlags().StringVarP(&edgeDeviceIP, "device-ip", "i", "10.0.0.1", "IP address to set. Set to \"0.0.0.0\" if you are using \"dhcp\" as --address-mode. If the edge is running the network's DHCP server this must be set explicitly; i.e. to \"192.168.1.0\" if the DHCP server should give out addresses from \"192.168.1.10\" to \"192.168.1.100\".")
	edgeCmd.PersistentFlags().StringVarP(&edgeDeviceNetmask, "device-netmask", "q", "255.255.255.0", "The netmask to use.")
	edgeCmd.PersistentFlags().StringVarP(&edgeDeviceMACAddress, "device-mac-address", "x", "DE:AD:BE:EF:01:10", "The MAC address to use. Must be unique per edge.")
	edgeCmd.PersistentFlags().IntVarP(&edgeMTU, "mtu", "u", 1290, "The MTU to use.")

	rootCmd.AddCommand(edgeCmd)
}
