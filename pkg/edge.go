package pkg

//go:generate sh -c "rm -rf n2n; git clone https://github.com/pojntfx/n2n.git; cd n2n; ./autogen.sh; ./configure; make"

/*
#cgo CFLAGS: -g3 -Wall
#cgo LDFLAGS: -lcrypto -ln2n -L${SRCDIR}/n2n
#include "edge.h"
*/
import "C"
import (
	"errors"
)

// Edge is a node which will be part of a virtual network.
type Edge struct {
	AllowP2P             bool   // Whether to allow peer-to-peer connections. If `false`, all traffic will be routed through the supernode.
	AllowRouting         bool   // Whether to allow the node to route traffic to other nodes.
	CommunityName        string // The name of the n2n community to join, i.e. `"myawesomecomunity"`.
	DisablePMTUDiscovery bool   // Whether to disable path MTU discovery.
	DisableMulticast     bool   // Whether to disable multicast.
	DynamicIPMode        bool   // Whether the IP address is set dynamically (see `AddressMode`). If the edge is running the network's DHCP server, this must be `false`.
	EncryptionKey        string // The key to use for encryption, i.e. `"mysecretkey"`.
	LocalPort            int    // The local port to use. `0` uses any available port.
	ManagementPort       int    // UDP management port. `5644` is the n2n default.
	RegisterInterval     int    // Interval in seconds for both UDP NAT hole punching and registration of the edge on the supernode. `1` is the n2n default.
	RegisterTTL          int    // Interval in seconds for UDP NAT hole punching through the supernode. `1` is the n2n default.
	SupernodeHostPort    string // Host:port of the supernode to connect to, i.e. `"localhost:1234"`.
	TypeOfService        int    // Type of service to use. `16` is the n2n default.
	EncryptionMethod     int    // Method of encryption to use. `1` is no encryption, `2` is Twofish encryption, `3` is AES-CBC encryption. Twofish encryption is the n2n default, but only due to legacy compatibility reasons; AES-CBC is up to ten times faster.
	DeviceName           string // Name of the TUN/TAP device to create, i.e. `"edge0"`.
	AddressMode          string // Mode of IP address assignment. `"static"` is a static assignment, `"dhcp"` uses the DHCP server at `DeviceIP` (see `DynamicIPMode`). If the edge is running the network's DHCP server, this must be `"static"`.
	DeviceIP             string // IP address to set. Set to `"0.0.0.0"` if you are using `"dhcp"` as `AddressMode`. If the edge is running the network's DHCP server this must be set explicitly; i.e. to `192.168.1.0` if the DHCP server should give out addresses from `192.168.1.10` to `192.168.1.100`.
	DeviceNetmask        string // The netmask to use, i.e. `"255.255.255.0"`.
	DeviceMACAddress     string // The MAC address to use, i.e. `"DE:AD:BE:EF:01:10"`. Must be unique per edge.
	MTU                  int    // The MTU to use. `1290` is the n2n default.
}

// getCIntFromGoBool converts a Go bool to a C int.
func (e *Edge) getCIntFromGoBool(goBool bool) C.int {
	cInt := C.int(0)
	if goBool {
		cInt = C.int(1)
	}

	return cInt
}

// Start starts an edge.
func (e *Edge) Start() error {
	res := int(C.edge_start(
		e.getCIntFromGoBool(e.AllowP2P),
		e.getCIntFromGoBool(e.AllowRouting),
		C.CString(e.CommunityName),
		e.getCIntFromGoBool(e.DisablePMTUDiscovery),
		e.getCIntFromGoBool(e.DisableMulticast),
		e.getCIntFromGoBool(e.DynamicIPMode),
		C.CString(e.EncryptionKey),
		C.int(e.LocalPort),
		C.int(e.ManagementPort),
		C.int(e.RegisterInterval),
		C.int(e.RegisterTTL),
		C.CString(e.SupernodeHostPort),
		C.int(e.TypeOfService),
		C.int(e.EncryptionMethod),
		C.CString(e.DeviceName),
		C.CString(e.AddressMode),
		C.CString(e.DeviceIP),
		C.CString(e.DeviceNetmask),
		C.CString(e.DeviceMACAddress),
		C.int(e.MTU)))

	if res == 0 {
		return nil
	}
	return errors.New("could not start edge")
}
