package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func listGroups(c grpc.ClientConnInterface, limit, offset int) (output *protos.ListGroupOutput, err error) {
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
func ListGroups(limit, offset int) (output *protos.ListGroupOutput, err error) {
	return listGroups(use().conn, limit, offset)
}

// ListGroups ...
func (cp *ConnProvider) ListGroups(limit, offset int) (output *protos.ListGroupOutput, err error) {
	return listGroups(cp.init().conn, limit, offset)
}

func listGroupsMap(c grpc.ClientConnInterface, limit, offset int) (output map[string]interface{}, err error) {
	output = make(map[string]interface{})
	groups, err := listGroups(c, limit, offset)
	if err != nil {
		return output, err
	}
	output = convert(groups.Data)
	output["count"] = groups.Count
	return output, nil
}

// ListGroupsMap ...
func ListGroupsMap(limit, offset int) (output map[string]interface{}, err error) {
	groups, err := listGroupsMap(use().conn, limit, offset)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

// ListGroupsMap ...
func (cp *ConnProvider) ListGroupsMap(limit, offset int) (output map[string]interface{}, err error) {
	groups, err := listGroupsMap(cp.init().conn, limit, offset)
	if err != nil {
		return nil, err
	}
	return groups, nil
}
