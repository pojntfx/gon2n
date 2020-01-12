package workers

/*
#cgo CFLAGS: -g3 -Wall
#cgo LDFLAGS: -lcrypto -ln2n -L${SRCDIR}/n2n
#include "edge.h"
*/
import "C"
import (
	"errors"
	"fmt"
	"strconv"
	"time"
	"unsafe"
)

// Edge is a node which will be part of a virtual network.
type Edge struct {
	AllowP2P             bool              // Whether to allow peer-to-peer connections. If `false`, all traffic will be routed through the supernode.
	AllowRouting         bool              // Whether to allow the node to route traffic to other nodes.
	CommunityName        string            // The name of the n2n community to join, i.e. `"myawesomecomunity"`.
	DisablePMTUDiscovery bool              // Whether to disable path MTU discovery.
	DisableMulticast     bool              // Whether to disable multicast.
	DynamicIPMode        bool              // Whether the IP address is set dynamically (see `AddressMode`). If the edge is running the network's DHCP server, this must be `false`.
	EncryptionKey        string            // The key to use for encryption, i.e. `"mysecretkey"`.
	LocalPort            int               // The local port to use. `0` uses any available port.
	ManagementPort       int               // UDP management port. `5644` is the n2n default.
	RegisterInterval     int               // Interval in seconds for both UDP NAT hole punching and registration of the edge on the supernode. `1` is the n2n default.
	RegisterTTL          int               // Interval in seconds for UDP NAT hole punching through the supernode. `1` is the n2n default.
	SupernodeHostPort    string            // Host:port of the supernode to connect to, i.e. `"localhost:1234"`.
	TypeOfService        int               // Type of service to use. `16` is the n2n default.
	EncryptionMethod     int               // Method of encryption to use. `1` is no encryption, `2` is Twofish encryption, `3` is AES-CBC encryption. Twofish encryption is the n2n default, but only due to legacy compatibility reasons; AES-CBC is up to ten times faster.
	DeviceName           string            // Name of the TUN/TAP device to create, i.e. `"edge0"`.
	AddressMode          string            // Mode of IP address assignment. `"static"` is a static assignment, `"dhcp"` uses the DHCP server at `DeviceIP` (see `DynamicIPMode`). If the edge is running the network's DHCP server, this must be `"static"`.
	DeviceIP             string            // IP address to set. Set to `"0.0.0.0"` if you are using `"dhcp"` as `AddressMode`. If the edge is running the network's DHCP server this must be set explicitly; i.e. to `192.168.1.0` if the DHCP server should give out addresses from `192.168.1.10` to `192.168.1.100`.
	DeviceNetmask        string            // The netmask to use, i.e. `"255.255.255.0"`.
	DeviceMACAddress     string            // The MAC address to use, i.e. `"DE:AD:BE:EF:01:10"`. Must be unique per edge.
	MTU                  int               // The MTU to use. `1290` is the n2n default, but `1500` is most common.
	cTuntapDevice        C.tuntap_dev      // TUN/TAP device instance.
	cConf                C.n2n_edge_conf_t // Internal edge configuration.
	cKeepRunning         C.int             // Whether the edge should be kept running. Set to `C.int(0)` at any time and it will be stopped.
}

// GetAllowP2P returns whether P2P is enabled.
func (e *Edge) GetAllowP2P() bool {
	varFromC, _ := e.getGoIntFromCUnsignedChar(e.cConf.allow_p2p)

	return e.getBoolFromInt(int(varFromC))
}

// GetAllowRouting returns whether routing is enabled.
func (e *Edge) GetAllowRouting() bool {
	varFromC, _ := e.getGoIntFromCUnsignedChar(e.cConf.allow_routing)

	return e.getBoolFromInt(int(varFromC))
}

// GetCommunityName returns the community name.
func (e *Edge) GetCommunityName() string {
	return C.GoString((*C.char)(unsafe.Pointer(&e.cConf.community_name)))
}

// GetDisablePMTUDiscovery returns whether path MTU discovery is enabled.
func (e *Edge) GetDisablePMTUDiscovery() bool {
	varFromC, _ := e.getGoIntFromCUnsignedChar(e.cConf.disable_pmtu_discovery)

	return e.getBoolFromInt(int(varFromC))
}

// GetDisableMulticast returns whether multicast is disabled.
func (e *Edge) GetDisableMulticast() bool {
	varFromC, _ := e.getGoIntFromCUnsignedChar(e.cConf.drop_multicast)

	return e.getBoolFromInt(int(varFromC))
}

// GetDynamicIPMode returns whether dynamic IP mode is enabled.
func (e *Edge) GetDynamicIPMode() bool {
	varFromC, _ := e.getGoIntFromCUnsignedChar(e.cConf.dyn_ip_mode)

	return e.getBoolFromInt(int(varFromC))
}

// GetEncryptionKey returns the encryption key.
func (e *Edge) GetEncryptionKey() string {
	return fmt.Sprintf("%T", e.cConf.encrypt_key)
}

// GetLocalPort returns the local port.
func (e *Edge) GetLocalPort() int {
	return int(e.cConf.local_port)
}

// GetManagementPort returns the management port.
func (e *Edge) GetManagementPort() int {
	return int(e.cConf.mgmt_port)
}

// GetRegisterInterval returns the register interval.
func (e *Edge) GetRegisterInterval() int {
	return int(e.cConf.register_interval)
}

// GetRegisterTTL returns the register TTL.
func (e *Edge) GetRegisterTTL() int {
	return int(e.cConf.register_ttl)
}

// GetTypeOfService returns the type of service.
func (e *Edge) GetTypeOfService() int {
	return int(e.cConf.tos)
}

// GetEncryptionMethod returns the encryption method.
func (e *Edge) GetEncryptionMethod() int {
	return int(e.cConf.transop_id)
}

// GetDeviceName returns the device name.
func (e *Edge) GetDeviceName() string {
	return C.GoString((*C.char)(unsafe.Pointer(&e.cTuntapDevice.dev_name)))
}

// GetMTU returns the MTU.
func (e *Edge) GetMTU() int {
	return int(e.cTuntapDevice.mtu)
}

// getGoIntFromCUnsignedChar parses the Go int from a C uchar.
func (e *Edge) getGoIntFromCUnsignedChar(cUChar C.uchar) (int64, error) {
	return strconv.ParseInt(fmt.Sprintf("%v", cUChar), 10, 64)
}

// getBoolFromInt converts the an int with the C conventions to a Go bool.
func (e *Edge) getBoolFromInt(cInt int) bool {
	return cInt != 0
}

// getCIntFromGoBool converts a Go bool to a C int.
func (e *Edge) getCIntFromGoBool(goBool bool) C.int {
	cInt := C.int(0)
	if goBool {
		cInt = C.int(1)
	}

	return cInt
}

// Configure configures the edge.
func (e *Edge) Configure() error {
	if errCode := C.edge_configure(
		&e.cConf,
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
		C.int(e.EncryptionMethod)); int(errCode) != 0 {
		return errors.New("could not configure edge" + string(e.RegisterInterval) + string(errCode))
	}

	return nil
}

// OpenTunTapDevice opens a TUN/TAP device for the edge.
func (e *Edge) OpenTunTapDevice() error {
	if errCode := C.tuntap_open(
		&e.cTuntapDevice,
		C.CString(e.DeviceName),
		C.CString(e.AddressMode),
		C.CString(e.DeviceIP),
		C.CString(e.DeviceNetmask),
		C.CString(e.DeviceMACAddress),
		C.int(e.MTU)); int(errCode) < 0 {
		return errors.New("could not open TUN/TAP device")
	}

	return nil
}

// Start starts an edge.
func (e *Edge) Start() error {
	e.cKeepRunning = C.int(1)

	if errCode := C.edge_start(
		&e.cTuntapDevice,
		&e.cConf,
		&e.cKeepRunning,
	); int(errCode) != 0 {
		return errors.New("could not start edge")
	}

	return nil
}

// Stop stops an edge.
func (e *Edge) Stop() error {
	e.cKeepRunning = C.int(0)

	return nil
}

// IsScheduledForDeletion returns true if the edge is scheduled for deletion.
func (e *Edge) IsScheduledForDeletion() bool {
	return int(e.cKeepRunning) == 0
}

// Wait blocks until the edge instance has stopped.
func (e *Edge) Wait() error {
	for {
		if int(e.cTuntapDevice.device_mask) == 0 {
			break
		}

		time.Sleep(1 * time.Second)
	}

	return nil
}
