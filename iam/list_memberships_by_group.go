package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func listMembershipsByGroup(c grpc.ClientConnInterface, input *protos.ListMembershipByGroupInput) (output *protos.ListMembershipJoinOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return protos.NewMembershipCRUDControllerClient(c).ListMembershipByGroup(ctx, input)
}

// ListMembershipsByGroup ...
func ListMembershipsByGroup(input *protos.ListMembershipByGroupInput) (output *protos.ListMembershipJoinOutput, err error) {
	return listMembershipsByGroup(use().conn, input)
}

// ListMembershipsByGroup ...
func (cp *ConnProvider) ListMembershipsByGroup(input *protos.ListMembershipByGroupInput) (output *protos.ListMembershipJoinOutput, err error) {
	return listMembershipsByGroup(cp.init().conn, input)
}

func listMembershipsByGroupMap(c grpc.ClientConnInterface, input *protos.ListMembershipByGroupInput) (output map[string]interface{}, err error) {
	output = make(map[string]interface{})
	memberships, err := listMembershipsByGroup(c, input)
	if err != nil {
		return output, err
	}
	output = convert(memberships.Data)
	output["count"] = memberships.Count
	return output, nil
}

// ListMembershipsByGroupMap ...
func ListMembershipsByGroupMap(input *protos.ListMembershipByGroupInput) (output map[string]interface{}, err error) {
	return listMembershipsByGroupMap(use().conn, input)
}

// ListMembershipsByGroupMap ...
func (cp *ConnProvider) ListMembershipsByGroupMap(input *protos.ListMembershipByGroupInput) (output map[string]interface{}, err error) {
	return listMembershipsByGroupMap(cp.init().conn, input)
}
