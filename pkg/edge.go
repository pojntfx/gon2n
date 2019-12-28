package pkg

//go:generate sh -c "rm -rf n2n; git clone https://github.com/ntop/n2n.git; cd n2n; ./autogen.sh; ./configure; make"

//#cgo CFLAGS: -g3 -Wall
//#cgo LDFLAGS: -lcrypto -ln2n -L${SRCDIR}/n2n
//#include "n2n/n2n.h"
import "C"

type Edge struct {
	DeviceName    string
	NetworkName   string
	SecretKey     string
	MyMacAddress  string
	MyIPv4Addr    string
	Supernode     string
	KeepOnRunning bool
}

func (e *Edge) Start() {
	deviceName := C.CString(e.DeviceName)
	networkName := C.CString(e.NetworkName)
	secretKey := C.CString(e.SecretKey)
	myMacAddress := C.CString(e.MyMacAddress)
	myIPv4Addr := C.CString(e.MyIPv4Addr)
	supernode := C.CString(e.Supernode)

	keepOnRunning := C.int(0)
	if !e.KeepOnRunning {
		keepOnRunning = C.int(1)
	}

	C.quick_edge_init(deviceName, networkName, secretKey, myMacAddress, myIPv4Addr, supernode, &keepOnRunning)
}