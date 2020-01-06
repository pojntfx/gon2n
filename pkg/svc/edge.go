package svc

import (
	"context"
	gon2n "github.com/pojntfx/gon2n/pkg/proto/generated"
	"github.com/pojntfx/gon2n/pkg/workers"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/bloom42/libs/rz-go/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// EdgeManager manages edges.
type EdgeManager struct {
	gon2n.UnimplementedEdgeManagerServer
	EdgesManaged map[string]*workers.Edge
}

// Create creates a edge.
func (e *EdgeManager) Create(_ context.Context, args *gon2n.EdgeManagerCreateArgs) (*gon2n.EdgeManagerCreateReply, error) {
	id := uuid.NewV4().String()

	edge := workers.Edge{
		AllowP2P:             args.GetAllowP2P(),
		AllowRouting:         args.GetAllowRouting(),
		CommunityName:        args.GetCommunityName(),
		DisablePMTUDiscovery: args.GetDisablePMTUDiscovery(),
		DisableMulticast:     args.GetDisableMulticast(),
		DynamicIPMode:        args.GetDynamicIPMode(),
		EncryptionKey:        args.GetEncryptionKey(),
		LocalPort:            int(args.GetLocalPort()),
		ManagementPort:       int(args.GetManagementPort()),
		RegisterInterval:     int(args.GetRegisterInterval()),
		RegisterTTL:          int(args.GetRegisterTTL()),
		SupernodeHostPort:    args.GetSupernodeHostPort(),
		TypeOfService:        int(args.GetTypeOfService()),
		EncryptionMethod:     int(args.GetEncryptionMethod()),
		DeviceName:           args.GetDeviceName(),
		AddressMode:          args.GetAddressMode(),
		DeviceIP:             args.GetDeviceIP(),
		DeviceNetmask:        args.GetDeviceNetmask(),
		DeviceMACAddress:     args.GetDeviceMACAddress(),
		MTU:                  int(args.GetMTU()),
	}

	if err := edge.Configure(); err != nil {
		log.Error(err.Error())

		return nil, status.Errorf(codes.Unknown, err.Error())
	}

	if err := edge.OpenTunTapDevice(); err != nil {
		log.Error(err.Error())

		return nil, status.Errorf(codes.Unknown, err.Error())
	}

	go func(edge *workers.Edge) {
		log.Info("Starting edge")

		if err := edge.Start(); err != nil {
			log.Error(err.Error())
		}

		// Keep the edge running
		for {
			if !edge.IsScheduledForDeletion() {
				log.Info("Restarting edge")

				if err := edge.Start(); err != nil {
					log.Error(err.Error())
				}
			} else {
				break
			}
		}
	}(&edge)

	e.EdgesManaged[id] = &edge

	return &gon2n.EdgeManagerCreateReply{
		Id: id,
	}, nil
}
