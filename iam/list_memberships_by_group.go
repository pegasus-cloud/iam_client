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

func listMembershipsByGroupMap(c grpc.ClientConnInterface, input *protos.ListMembershipByGroupInput) (output map[string]*protos.MemberJoin, count int64, err error) {
	memberships, err := listMembershipsByGroup(c, input)
	for _, membership := range memberships.Data {
		output[membership.ID] = membership
	}
	return output, memberships.Count, err
}

// ListMembershipsByGroupMap ...
func ListMembershipsByGroupMap(input *protos.ListMembershipByGroupInput) (output map[string]*protos.MemberJoin, count int64, err error) {
	return listMembershipsByGroupMap(use().conn, input)
}

// ListMembershipsByGroupMap ...
func (cp *ConnProvider) ListMembershipsByGroupMap(input *protos.ListMembershipByGroupInput) (output map[string]*protos.MemberJoin, count int64, err error) {
	return listMembershipsByGroupMap(cp.init().conn, input)
}
