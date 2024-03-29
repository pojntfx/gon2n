package services

import (
	"context"

	api "github.com/pojntfx/gon2n/pkg/api/proto/v1"
	"github.com/pojntfx/gon2n/pkg/workers"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/bloom42/libs/rz-go/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// EdgeManager manages edges.
type EdgeManager struct {
	api.UnimplementedEdgeManagerServer
	EdgesManaged map[string]*workers.Edge
}

// Create creates a edge.
func (e *EdgeManager) Create(_ context.Context, args *api.Edge) (*api.EdgeManagedId, error) {
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

	return &api.EdgeManagedId{
		Id: id,
	}, nil
}

// List lists the managed edges.
func (e *EdgeManager) List(_ context.Context, args *api.EdgeManagerListArgs) (*api.EdgeManagerListReply, error) {
	log.Info("Listing edges")

	var edgesManaged []*api.EdgeManaged
	for id, edge := range e.EdgesManaged {
		edgeManaged := api.EdgeManaged{
			Id:                   id,
			AllowP2P:             edge.GetAllowP2P(),
			AllowRouting:         edge.GetAllowRouting(),
			CommunityName:        edge.GetCommunityName(),
			DisablePMTUDiscovery: edge.GetDisablePMTUDiscovery(),
			DisableMulticast:     edge.GetDisableMulticast(),
			DynamicIPMode:        edge.GetDynamicIPMode(),
			LocalPort:            int64(edge.GetLocalPort()),
			ManagementPort:       int64(edge.GetManagementPort()),
			RegisterInterval:     int64(edge.GetRegisterInterval()),
			RegisterTTL:          int64(edge.GetRegisterTTL()),
			SupernodeHostPort:    edge.SupernodeHostPort,
			TypeOfService:        int64(edge.GetTypeOfService()),
			EncryptionMethod:     int64(edge.GetEncryptionMethod()),
			DeviceName:           edge.GetDeviceName(),
			AddressMode:          edge.AddressMode,
			DeviceIP:             edge.DeviceIP,
			DeviceNetmask:        edge.DeviceNetmask,
			DeviceMACAddress:     edge.DeviceMACAddress,
			MTU:                  int64(edge.GetMTU()),
		}

		edgesManaged = append(edgesManaged, &edgeManaged)
	}

	return &api.EdgeManagerListReply{
		EdgesManaged: edgesManaged,
	}, nil
}

// Get gets one of the managed edges.
func (e *EdgeManager) Get(_ context.Context, args *api.EdgeManagedId) (*api.EdgeManaged, error) {
	log.Info("Getting edge")

	var edgeManaged *api.EdgeManaged

	for id, edge := range e.EdgesManaged {
		if id == args.GetId() {
			edgeManaged = &api.EdgeManaged{
				Id:                   id,
				AllowP2P:             edge.GetAllowP2P(),
				AllowRouting:         edge.GetAllowRouting(),
				CommunityName:        edge.GetCommunityName(),
				DisablePMTUDiscovery: edge.GetDisablePMTUDiscovery(),
				DisableMulticast:     edge.GetDisableMulticast(),
				DynamicIPMode:        edge.GetDynamicIPMode(),
				LocalPort:            int64(edge.GetLocalPort()),
				ManagementPort:       int64(edge.GetManagementPort()),
				RegisterInterval:     int64(edge.GetRegisterInterval()),
				RegisterTTL:          int64(edge.GetRegisterTTL()),
				SupernodeHostPort:    edge.SupernodeHostPort,
				TypeOfService:        int64(edge.GetTypeOfService()),
				EncryptionMethod:     int64(edge.GetEncryptionMethod()),
				DeviceName:           edge.GetDeviceName(),
				AddressMode:          edge.AddressMode,
				DeviceIP:             edge.DeviceIP,
				DeviceNetmask:        edge.DeviceNetmask,
				DeviceMACAddress:     edge.DeviceMACAddress,
				MTU:                  int64(edge.GetMTU()),
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
func (e *EdgeManager) Delete(_ context.Context, args *api.EdgeManagedId) (*api.EdgeManagedId, error) {
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

	return &api.EdgeManagedId{
		Id: id,
	}, nil
}
