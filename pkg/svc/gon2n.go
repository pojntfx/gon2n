package svc

//go:generate sh -c "mkdir -p ../proto/generated && protoc --go_out=paths=source_relative,plugins=grpc:../proto/generated -I=../ ../proto/*.proto"

import (
	"context"
	gon2n "github.com/pojntfx/gon2n/pkg/proto/generated/proto"
	"gitlab.com/bloom42/libs/rz-go/v2/log"
)

// SupernodeManager manages supernodes.
type SupernodeManager struct {
	gon2n.UnimplementedSupernodeManagerServer
}

// Create creates a supernode.
func (s *SupernodeManager) Create(_ context.Context, args *gon2n.SupernodeManagerCreateArgs) (*gon2n.SupernodeManagerCreateReply, error) {
	log.Info("Creating supernode")

	return &gon2n.SupernodeManagerCreateReply{
		Id: 1,
	}, nil
}
