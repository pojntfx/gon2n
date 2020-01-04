package svc

//go:generate sh -c "mkdir -p ../proto/generated && protoc --go_out=paths=source_relative,plugins=grpc:../proto/generated -I=../ ../proto/*.proto"

import (
	"context"
	gon2n "github.com/pojntfx/gon2n/pkg/proto/generated/proto"
	"github.com/pojntfx/gon2n/pkg/workers"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/bloom42/libs/rz-go/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SupernodeManager manages supernodes.
type SupernodeManager struct {
	gon2n.UnimplementedSupernodeManagerServer
	SupernodesManaged map[string]*workers.Supernode
}

// Create creates a supernode.
func (s *SupernodeManager) Create(_ context.Context, args *gon2n.SupernodeManagerCreateArgs) (*gon2n.SupernodeManagerCreateReply, error) {
	id := uuid.NewV4().String()

	supernode := workers.Supernode{
		ListenPort:     int(args.GetListenPort()),
		ManagementPort: int(args.GetManagementPort()),
	}

	if err := supernode.Configure(); err != nil {
		log.Error(err.Error())

		return nil, status.Errorf(codes.Unknown, err.Error())
	}

	if err := supernode.OpenListenPortSocket(); err != nil {
		log.Error(err.Error())

		return nil, status.Errorf(codes.Unknown, err.Error())
	}

	if err := supernode.OpenManagementPortSocket(); err != nil {
		log.Error(err.Error())

		return nil, status.Errorf(codes.Unknown, err.Error())
	}

	log.Info("Starting supernode")

	go func(supernode *workers.Supernode) {
		// Keep the supernode running
		for {
			if err := supernode.Start(); err != nil {
				log.Error(err.Error())
			}
		}
	}(&supernode)

	s.SupernodesManaged[id] = &supernode

	return &gon2n.SupernodeManagerCreateReply{
		Id: id,
	}, nil
}

// List lists the managed supernodes.
func (s *SupernodeManager) List(_ context.Context, args *gon2n.SupernodeManagerListArgs) (*gon2n.SupernodeManagerListReply, error) {
	var supernodesManaged []*gon2n.SupernodeManaged

	for id := range s.SupernodesManaged {
		supernodeManaged := gon2n.SupernodeManaged{
			Id: id,
		}

		supernodesManaged = append(supernodesManaged, &supernodeManaged)
	}

	return &gon2n.SupernodeManagerListReply{
		SupernodesManaged: supernodesManaged,
	}, nil
}
