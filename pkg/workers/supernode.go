package workers

//go:generate sh -c "rm -rf n2n; git clone https://github.com/ntop/n2n.git; cd n2n; git checkout 2.8-stable; ./autogen.sh; ./configure --host=${CC}; make -j $(nproc); ar dvf libn2n.a edge.o; ar dvf libn2n.a example_edge_embed.o; ar dvf libn2n.a example_edge_embed_quick_edge_init.o; ar dvf libn2n.a example_sn_embed.o; ar dvf libn2n.a sn.o"

/*
#cgo CFLAGS: -g3 -Wall
#cgo LDFLAGS: -lcrypto -ln2n -L${SRCDIR}/n2n
#include "supernode.h"
*/
import "C"
import (
	"errors"
	"time"
)

// Supernode allows edge nodes to announce and discover other nodes.
type Supernode struct {
	ListenPort     int        // UDP listen port.
	ManagementPort int        // UDP management port. `5645` is the n2n default.
	cSupernode     C.n2n_sn_t // Supernode instance.
	cKeepRunning   C.int      // Whether the supernode should be kept running. Set to `C.int(0)` at any time and it will be stopped.
}

// Configure configures a supernode.
func (s *Supernode) Configure() error {
	if errCode := C.supernode_configure(&s.cSupernode, C.int(s.ListenPort)); int(errCode) != 0 {
		return errors.New("could not configure supernode")
	}

	return nil
}

// OpenListenPortSocket opens a listen port socket.
func (s *Supernode) OpenListenPortSocket() error {
	if errCode := C.supernode_open_lport_socket(&s.cSupernode); int(errCode) != 0 {
		return errors.New("could not open listen port socket")
	}

	return nil
}

// OpenManagementPortSocket opens a management port socket.
func (s *Supernode) OpenManagementPortSocket() error {
	if errCode := C.supernode_open_mgmt_socket(&s.cSupernode, C.int(s.ManagementPort)); int(errCode) != 0 {
		return errors.New("could not open management port socket")
	}

	return nil
}

// Start starts a supernode.
func (s *Supernode) Start() error {
	s.cKeepRunning = C.int(1)

	if errCode := C.supernode_start(&s.cSupernode, &s.cKeepRunning); int(errCode) != 0 {
		return errors.New("could not start supernode")
	}

	return nil
}

// Stop stops a supernode.
func (s *Supernode) Stop() error {
	s.cKeepRunning = C.int(0)

	return nil
}

// IsScheduledForDeletion returns true if the supernode is scheduled for deletion.
func (s *Supernode) IsScheduledForDeletion() bool {
	return int(s.cKeepRunning) == 0
}

// Wait blocks until the supernode instance has stopped.
func (s *Supernode) Wait() error {
	for {
		if int(s.cSupernode.mgmt_sock) == -1 {
			break
		}

		time.Sleep(1 * time.Second)
	}

	return nil
}

// GetListenPort returns the listen port of the supernode.
func (s *Supernode) GetListenPort() int {
	return int(C.int(s.cSupernode.lport))
}
