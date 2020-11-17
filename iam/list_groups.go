package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func listGroups(c grpc.ClientConnInterface, limit, offset int) (output *protos.GroupInfos, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	groups, err := protos.NewGroupCURDControllerClient(c).ListGroup(ctx, &protos.LimitOffset{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}
	return groups, nil
}

// ListGroups ...
func ListGroups(limit, offset int) (output *protos.GroupInfos, err error) {
	return listGroups(use().conn, limit, offset)
}

// ListGroups ...
func (cp *ConnProvider) ListGroups(limit, offset int) (output *protos.GroupInfos, err error) {
	return listGroups(cp.init().conn, limit, offset)
}
