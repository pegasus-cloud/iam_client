package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func listMembershipsByGroup(c grpc.ClientConnInterface, groupID string, limit, offset int) (output *protos.ListMembershipJoinOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	memberships, err := protos.NewMembershipCRUDControllerClient(c).ListMembershipByGroup(ctx, &protos.ListMembershipByGroupInput{
		GroupID: groupID,
		Data: &protos.LimitOffset{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	})
	if err != nil {
		return nil, err
	}
	return memberships, nil
}

// ListMembershipsByGroup ...
func ListMembershipsByGroup(groupID string, limit, offset int) (output *protos.ListMembershipJoinOutput, err error) {
	return listMembershipsByGroup(use().conn, groupID, limit, offset)
}

// ListMembershipsByGroup ...
func (cp *ConnProvider) ListMembershipsByGroup(groupID string, limit, offset int) (output *protos.ListMembershipJoinOutput, err error) {
	return listMembershipsByGroup(cp.init().conn, groupID, limit, offset)
}

func listMembershipsByGroupMap(c grpc.ClientConnInterface, groupID string, limit, offset int) (output map[string]interface{}, err error) {
	output = make(map[string]interface{})
	memberships, err := listMembershipsByGroup(c, groupID, limit, offset)
	if err != nil {
		return output, err
	}
	output = convert(memberships.Data)
	output["count"] = memberships.Count
	return output, nil
}

// ListMembershipsByGroupMap ...
func ListMembershipsByGroupMap(groupID string, limit, offset int) (output map[string]interface{}, err error) {
	memberships, err := listMembershipsByGroupMap(use().conn, groupID, limit, offset)
	if err != nil {
		return nil, err
	}
	return memberships, nil
}

// ListMembershipsByGroupMap ...
func (cp *ConnProvider) ListMembershipsByGroupMap(groupID string, limit, offset int) (output map[string]interface{}, err error) {
	memberships, err := listMembershipsByGroupMap(cp.init().conn, groupID, limit, offset)
	if err != nil {
		return nil, err
	}
	return memberships, nil
}
