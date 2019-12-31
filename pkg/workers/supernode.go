package workers

//go:generate sh -c "rm -rf n2n; git clone https://github.com/pojntfx/n2n.git; cd n2n; ./autogen.sh; ./configure; make"

/*
#cgo CFLAGS: -g3 -Wall
#cgo LDFLAGS: -lcrypto -ln2n -L${SRCDIR}/n2n
#include "supernode.h"
*/
import "C"
import (
	"errors"
)

// Supernode allows edge nodes to announce and discover other nodes.
type Supernode struct {
	ListenPort     int        // UDP listen port.
	ManagementPort int        // UDP management port. `5645` is the n2n default.
	cSupernode     C.n2n_sn_t // Supernode instance.
	cKeepRunning   C.int      // Whether the supernode should be kept running. Set to `C.int(0)` at any time and it will be stopped.
}

// Configure configures a supernode.
func (e *Supernode) Configure() error {
	if errCode := C.supernode_configure(&e.cSupernode, C.int(e.ListenPort)); int(errCode) != 0 {
		return errors.New("could not configure supernode")
	}

	return nil
}

// OpenListenPortSocket opens a listen port socket.
func (e *Supernode) OpenListenPortSocket() error {
	if errCode := C.supernode_open_lport_socket(&e.cSupernode); int(errCode) != 0 {
		return errors.New("could not open listen port socket")
	}

	return nil
}

// OpenManagementPortSocket opens a management port socket.
func (e *Supernode) OpenManagementPortSocket() error {
	if errCode := C.supernode_open_mgmt_socket(&e.cSupernode, C.int(e.ManagementPort)); int(errCode) != 0 {
		return errors.New("could not open management port socket")
	}

	return nil
}

// Start starts a supernode.
func (e *Supernode) Start() error {
	e.cKeepRunning = C.int(1)

	if errCode := C.supernode_start(&e.cSupernode, &e.cKeepRunning); int(errCode) != 0 {
		return errors.New("could not start supernode")
	}

	return nil
}

// Stop stops a supernode.
func (e *Supernode) Stop() error {
	e.cKeepRunning = C.int(0)

	return nil
}
