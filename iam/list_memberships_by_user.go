package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func listMembershipsByUser(c grpc.ClientConnInterface, input *protos.ListMembershipByUserInput) (output *protos.ListMembershipJoinOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return protos.NewMembershipCRUDControllerClient(c).ListMembershipByUser(ctx, input)
}

// ListMembershipsByUser ...
func ListMembershipsByUser(input *protos.ListMembershipByUserInput) (output *protos.ListMembershipJoinOutput, err error) {
	return listMembershipsByUser(use().conn, input)
}

// ListMembershipsByUser ...
func (cp *ConnProvider) ListMembershipsByUser(input *protos.ListMembershipByUserInput) (output *protos.ListMembershipJoinOutput, err error) {
	return listMembershipsByUser(cp.init().conn, input)
}

func listMembershipsByUserMap(c grpc.ClientConnInterface, input *protos.ListMembershipByUserInput) (output map[string]*protos.MemberJoin, count int64, err error) {
	memberships, err := listMembershipsByUser(c, input)
	for _, membership := range memberships.Data {
		output[membership.ID] = membership
	}
	return output, memberships.Count, err
}

// ListMembershipsByUserMap ...
func ListMembershipsByUserMap(input *protos.ListMembershipByUserInput) (output map[string]*protos.MemberJoin, count int64, err error) {
	return listMembershipsByUserMap(use().conn, input)
}

// ListMembershipsByUserMap ...
func (cp *ConnProvider) ListMembershipsByUserMap(input *protos.ListMembershipByUserInput) (output map[string]*protos.MemberJoin, count int64, err error) {
	return listMembershipsByUserMap(cp.init().conn, input)
}
