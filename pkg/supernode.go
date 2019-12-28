package pkg

//go:generate sh -c "rm -rf n2n; git clone https://github.com/ntop/n2n.git; cd n2n; ./autogen.sh; ./configure; make"

/*
#cgo CFLAGS: -g3 -Wall
#cgo LDFLAGS: -lcrypto -ln2n -L${SRCDIR}/n2n
#include "supernode.h"
*/
import "C"

type Supernode struct {
}

func (e *Supernode) Start() {
	C.supernode_start()
}
