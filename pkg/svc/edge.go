package svc

import (
	gon2n "github.com/pojntfx/gon2n/pkg/proto/generated"
	"github.com/pojntfx/gon2n/pkg/workers"
)

// EdgeManager manages edges.
type EdgeManager struct {
	gon2n.UnimplementedEdgeManagerServer
	EdgesManaged map[string]*workers.Edge
}
