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
}

func (e *Edge) Start() error {
	res := int(C.edge_start())

	if res == 0 {
		return nil
	}
	return errors.New("could not start edge")
}
