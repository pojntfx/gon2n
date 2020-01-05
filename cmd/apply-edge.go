package cmd

import (
	"github.com/pojntfx/gon2n/pkg/workers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bloom42/libs/rz-go/v2"
	"gitlab.com/bloom42/libs/rz-go/v2/log"
	"os"
	"os/signal"
	"syscall"
)

var applyEdgeCmd = &cobra.Command{
	Use:     "edge",
	Aliases: []string{"edges", "e"},
	Short:   "Apply an edge",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !(viper.GetString(edgeConfigFileKey) == edgeConfigFileDefault) {
			viper.SetConfigFile(viper.GetString(edgeConfigFileKey))

			if err := viper.ReadInConfig(); err != nil {
				return err
			}
		}

		edge := workers.Edge{
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

		if err := edge.Configure(); err != nil {
			return err
		}

		if err := edge.OpenTunTapDevice(); err != nil {
			return err
		}

		interrupt := make(chan os.Signal, 2)
		signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-interrupt

			// Allow manually killing the process
			go func() {
				<-interrupt

				os.Exit(1)
			}()

			log.Info("Gracefully stopping edge (this might take a few seconds)")

			if err := edge.Stop(); err != nil {
				log.Fatal("Could not stop edge", rz.Err(err))
			}
		}()

		log.Info("Starting edge")

		return edge.Start()
	},
}

func init() {
	var (
		edgeConfigFileFlag           string
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

	applyEdgeCmd.PersistentFlags().StringVarP(&edgeConfigFileFlag, edgeConfigFileKey, "f", edgeConfigFileDefault, "Configuration file to use.")
	applyEdgeCmd.PersistentFlags().BoolVarP(&edgeAllowP2PFlag, edgeAllowP2PKey, "p", true, "Whether to allow peer-to-peer connections. If false, all traffic will be routed through the supernode.")
	applyEdgeCmd.PersistentFlags().BoolVarP(&edgeAllowRoutingFlag, edgeAllowRoutingKey, "r", true, "Whether to allow the node to route traffic to other nodes.")
	applyEdgeCmd.PersistentFlags().StringVarP(&edgeCommunityNameFlag, edgeCommunityNameKey, "c", "mynetwork", "The name of the n2n community to join.")
	applyEdgeCmd.PersistentFlags().BoolVarP(&edgeDisablePMTUDiscoveryFlag, edgeDisablePMTUDiscoveryKey, "d", false, "Whether to disable path MTU discovery.")
	applyEdgeCmd.PersistentFlags().BoolVarP(&edgeDisableMulticastFlag, edgeDisableMulticastKey, "m", false, "Whether to disable multicast.")
	applyEdgeCmd.PersistentFlags().BoolVarP(&edgeDynamicIPModeFlag, edgeDynamicIPModeKey, "y", false, "Whether the IP address is set dynamically (see --address-mode). If the edge is running the network's DHCP server, this must be false.")
	applyEdgeCmd.PersistentFlags().StringVarP(&edgeEncryptionKeyFlag, edgeEncryptionKeyKey, "k", "mysecretkey", "The key to use for encryption.")
	applyEdgeCmd.PersistentFlags().IntVarP(&edgeLocalPortFlag, edgeLocalPortKey, "l", 0, "The local port to use. 0 uses any available port.")
	applyEdgeCmd.PersistentFlags().IntVarP(&edgeManagementPortFlag, edgeManagementPortKey, "a", 5644, "UDP management port. 5644 is the n2n default.")
	applyEdgeCmd.PersistentFlags().IntVarP(&edgeRegisterIntervalFlag, edgeRegisterIntervalKey, "n", 1, "Interval in seconds for both UDP NAT hole punching and registration of the edge on the supernode.")
	applyEdgeCmd.PersistentFlags().IntVarP(&edgeRegisterTTLFlag, edgeRegisterTTLKey, "t", 1, "Interval in seconds for UDP NAT hole punching through the supernode.")
	applyEdgeCmd.PersistentFlags().StringVarP(&edgeSupernodeHostPortFlag, edgeSupernodeHostPortKey, "s", "localhost:1234", "Host:port of the supernode to connect to.")
	applyEdgeCmd.PersistentFlags().IntVarP(&edgeTypeOfServiceFlag, edgeTypeOfServiceKey, "o", 16, "Type of service to use.")
	applyEdgeCmd.PersistentFlags().IntVarP(&edgeEncryptionMethodFlag, edgeEncryptionMethodKey, "e", 2, "Method of encryption to use. 1 is no encryption, 2 is Twofish encryption, 3 is AES-CBC encryption. Twofish encryption is the n2n default, but only due to legacy compatibility reasons; AES-CBC is up to ten times faster.")
	applyEdgeCmd.PersistentFlags().StringVarP(&edgeDeviceNameFlag, edgeDeviceNameKey, "v", "edge0", "Name of the TUN/TAP device to create.")
	applyEdgeCmd.PersistentFlags().StringVarP(&edgeAddressModeFlag, edgeAddressModeKey, "z", "static", "Mode of IP address assignment. \"static\" is a static assignment, \"dhcp\" uses the DHCP server at --device-ip (see --dynamic-ip-node). If the edge is running the network's DHCP server, this must be \"static\".")
	applyEdgeCmd.PersistentFlags().StringVarP(&edgeDeviceIPFlag, edgeDeviceIPKey, "i", "10.0.0.1", "IP address to set. Set to \"0.0.0.0\" if you are using \"dhcp\" as --address-mode. If the edge is running the network's DHCP server this must be set explicitly; i.e. to \"192.168.1.0\" if the DHCP server should give out addresses from \"192.168.1.10\" to \"192.168.1.100\".")
	applyEdgeCmd.PersistentFlags().StringVarP(&edgeDeviceNetmaskFlag, edgeDeviceNetmaskKey, "q", "255.255.255.0", "The netmask to use.")
	applyEdgeCmd.PersistentFlags().StringVarP(&edgeDeviceMACAddressFlag, edgeDeviceMACAddressKey, "x", "DE:AD:BE:EF:01:10", "The MAC address to use. Must be unique per edge.")
	applyEdgeCmd.PersistentFlags().IntVarP(&edgeMTUFlag, edgeMTUKey, "u", 1290, "The MTU to use.")

	if err := viper.BindPFlags(applyEdgeCmd.PersistentFlags()); err != nil {
		log.Fatal(couldNotBindFlagsErrorMessage, rz.Err(err))
	}

	viper.AutomaticEnv()

	applyCmd.AddCommand(applyEdgeCmd)
}
