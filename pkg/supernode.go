package pkg

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

type Supernode struct {
	ListenPort int;
	ManagementPort int;
}

func (e *Supernode) Start() error {
	res := int(C.supernode_start(C.int(e.ListenPort), C.int(e.ManagementPort)))

	if res == 0 {
		return nil
	}
	return errors.New("could not start supernode")
}
