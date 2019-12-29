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

type Edge struct {
	AllowP2P             bool
	AllowRouting         bool
	CommunityName        string
	DisablePMTUDiscovery bool
	DropMulticast        bool
	DynamicIPMode        bool
	EncryptionKey        string
	LocalPort            int
	ManagementPort       int
	RegisterInterval     int
	RegisterTTL          int
	SupernodeHostPort    string
	TypeOfService        int
	EncryptionMethod     int
	DeviceName           string
	AddressMode          string
	DeviceIP             string
	DeviceNetmask        string
	DeviceMACAddress     string
	MTU                  int
}

func (e *Edge) getCIntFromGoBool(goBool bool) C.int {
	cInt := C.int(0)
	if goBool {
		cInt = C.int(1)
	}

	return cInt
}

func (e *Edge) Start() error {
	res := int(C.edge_start(
		e.getCIntFromGoBool(e.AllowP2P),
		e.getCIntFromGoBool(e.AllowRouting),
		C.CString(e.CommunityName),
		e.getCIntFromGoBool(e.DisablePMTUDiscovery),
		e.getCIntFromGoBool(e.DropMulticast),
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
