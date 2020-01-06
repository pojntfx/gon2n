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

// List lists the managed edges.
func (e *EdgeManager) List(_ context.Context, args *gon2n.EdgeManagerListArgs) (*gon2n.EdgeManagerListReply, error) {
	log.Info("Listing edges")

	var edgesManaged []*gon2n.EdgeManaged

	for id, edge := range e.EdgesManaged {
		edgeManaged := gon2n.EdgeManaged{
			Id:                   id,
			AllowP2P:             edge.AllowP2P,
			AllowRouting:         edge.AllowRouting,
			CommunityName:        edge.CommunityName,
			DisablePMTUDiscovery: edge.DisablePMTUDiscovery,
			DisableMulticast:     edge.DisableMulticast,
			DynamicIPMode:        edge.DynamicIPMode,
			LocalPort:            int64(edge.LocalPort),
			ManagementPort:       int64(edge.ManagementPort),
			RegisterInterval:     int64(edge.RegisterInterval),
			RegisterTTL:          int64(edge.RegisterTTL),
			SupernodeHostPort:    edge.SupernodeHostPort,
			TypeOfService:        int64(edge.TypeOfService),
			EncryptionMethod:     int64(edge.EncryptionMethod),
			DeviceName:           edge.DeviceName,
			AddressMode:          edge.AddressMode,
			DeviceIP:             edge.DeviceIP,
			DeviceNetmask:        edge.DeviceNetmask,
			DeviceMACAddress:     edge.DeviceMACAddress,
			MTU:                  int64(edge.MTU),
		}

		edgesManaged = append(edgesManaged, &edgeManaged)
	}

	return &gon2n.EdgeManagerListReply{
		EdgesManaged: edgesManaged,
	}, nil
}

// Get gets one of the managed edges.
func (e *EdgeManager) Get(_ context.Context, args *gon2n.EdgeManagerGetArgs) (*gon2n.EdgeManaged, error) {
	log.Info("Getting edge")

	var edgeManaged *gon2n.EdgeManaged

	for id, edge := range e.EdgesManaged {
		if id == args.GetId() {
			edgeManaged = &gon2n.EdgeManaged{
				Id:                   id,
				AllowP2P:             edge.AllowP2P,
				AllowRouting:         edge.AllowRouting,
				CommunityName:        edge.CommunityName,
				DisablePMTUDiscovery: edge.DisablePMTUDiscovery,
				DisableMulticast:     edge.DisableMulticast,
				DynamicIPMode:        edge.DynamicIPMode,
				LocalPort:            int64(edge.LocalPort),
				ManagementPort:       int64(edge.ManagementPort),
				RegisterInterval:     int64(edge.RegisterInterval),
				RegisterTTL:          int64(edge.RegisterTTL),
				SupernodeHostPort:    edge.SupernodeHostPort,
				TypeOfService:        int64(edge.TypeOfService),
				EncryptionMethod:     int64(edge.EncryptionMethod),
				DeviceName:           edge.DeviceName,
				AddressMode:          edge.AddressMode,
				DeviceIP:             edge.DeviceIP,
				DeviceNetmask:        edge.DeviceNetmask,
				DeviceMACAddress:     edge.DeviceMACAddress,
				MTU:                  int64(edge.MTU),
			}
			break
		}
	}

	if edgeManaged != nil {
		return edgeManaged, nil
	}

	msg := "edge not found"

	log.Error(msg)

	return nil, status.Errorf(codes.NotFound, msg)
}

// Delete deletes a edge.
func (e *EdgeManager) Delete(_ context.Context, args *gon2n.EdgeManagerDeleteArgs) (*gon2n.EdgeManagerDeleteReply, error) {
	id := args.GetId()

	edgesManaged := e.EdgesManaged[id]
	if edgesManaged == nil {
		msg := "edge not found"

		log.Error(msg)

		return nil, status.Errorf(codes.NotFound, msg)
	}

	log.Info("Stopping edge")

	if err := edgesManaged.Stop(); err != nil {
		log.Error(err.Error())

		return nil, status.Errorf(codes.Unknown, err.Error())
	}

	if err := edgesManaged.Wait(); err != nil {
		log.Error(err.Error())

		return nil, status.Errorf(codes.Unknown, err.Error())
	}

	delete(e.EdgesManaged, id)

	return &gon2n.EdgeManagerDeleteReply{
		Id: id,
	}, nil
}
