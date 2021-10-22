package services

//go:generate sh -c "mkdir -p ../api/proto/v1 && protoc --go_out=paths=source_relative,plugins=grpc:../api/proto/v1 -I=../../api/proto/v1 ../../api/proto/v1/*.proto"

import (
	"context"

	api "github.com/pojntfx/gon2n/pkg/api/proto/v1"
	"github.com/pojntfx/gon2n/pkg/workers"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/bloom42/libs/rz-go/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SupernodeManager manages supernodes.
type SupernodeManager struct {
	api.UnimplementedSupernodeManagerServer
	SupernodesManaged map[string]*workers.Supernode
}

// Create creates a supernode.
func (s *SupernodeManager) Create(_ context.Context, args *api.Supernode) (*api.SupernodeManagedId, error) {
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

	go func(supernode *workers.Supernode) {
		log.Info("Starting supernode")

		if err := supernode.Start(); err != nil {
			log.Error(err.Error())
		}

		// Keep the supernode running
		for {
			if !supernode.IsScheduledForDeletion() {
				log.Info("Restarting supernode")

				if err := supernode.Start(); err != nil {
					log.Error(err.Error())
				}
			} else {
				break
			}
		}
	}(&supernode)

	s.SupernodesManaged[id] = &supernode

	return &api.SupernodeManagedId{
		Id: id,
	}, nil
}

// List lists the managed supernodes.
func (s *SupernodeManager) List(_ context.Context, args *api.SupernodeManagerListArgs) (*api.SupernodeManagerListReply, error) {
	log.Info("Listing supernodes")

	var supernodesManaged []*api.SupernodeManaged
	for id, supernode := range s.SupernodesManaged {
		supernodeManaged := api.SupernodeManaged{
			Id:             id,
			ListenPort:     int64(supernode.GetListenPort()),
			ManagementPort: int64(supernode.ManagementPort),
		}

		supernodesManaged = append(supernodesManaged, &supernodeManaged)
	}

	return &api.SupernodeManagerListReply{
		SupernodesManaged: supernodesManaged,
	}, nil
}

// Get gets one of the managed supernodes.
func (s *SupernodeManager) Get(_ context.Context, args *api.SupernodeManagedId) (*api.SupernodeManaged, error) {
	log.Info("Getting supernode")

	var supernodeManaged *api.SupernodeManaged

	for id, supernode := range s.SupernodesManaged {
		if id == args.GetId() {
			supernodeManaged = &api.SupernodeManaged{
				Id:             id,
				ListenPort:     int64(supernode.GetListenPort()),
				ManagementPort: int64(supernode.ManagementPort),
			}
			break
		}
	}

	if supernodeManaged != nil {
		return supernodeManaged, nil
	}

	msg := "supernode not found"

	log.Error(msg)

	return nil, status.Errorf(codes.NotFound, msg)
}

// Delete deletes a supernode.
func (s *SupernodeManager) Delete(_ context.Context, args *api.SupernodeManagedId) (*api.SupernodeManagedId, error) {
	id := args.GetId()

	supernodesManaged := s.SupernodesManaged[id]
	if supernodesManaged != nil {
		log.Info("Stopping supernode")

		if err := supernodesManaged.Stop(); err != nil {
			log.Error(err.Error())

			return nil, status.Errorf(codes.Unknown, err.Error())
		}

		if err := supernodesManaged.Wait(); err != nil {
			log.Error(err.Error())

			return nil, status.Errorf(codes.Unknown, err.Error())
		}

		delete(s.SupernodesManaged, id)

		return &api.SupernodeManagedId{
			Id: id,
		}, nil
	}

	msg := "supernode not found"

	log.Error(msg)

	return nil, status.Errorf(codes.NotFound, msg)
}
