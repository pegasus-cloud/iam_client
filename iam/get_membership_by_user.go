package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func getMembershipByUser(c grpc.ClientConnInterface, input *protos.MemUserGroupInput) (output *protos.MemberJoin, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return protos.NewMembershipCRUDControllerClient(c).GetMembershipJoin(ctx, input)
}

// GetMembershipByUser ...
func GetMembershipByUser(input *protos.MemUserGroupInput) (output *protos.MemberJoin, err error) {
	return getMembershipByUser(use().conn, input)
}

// GetMembershipByUser ...
func (cp *ConnProvider) GetMembershipByUser(input *protos.MemUserGroupInput) (output *protos.MemberJoin, err error) {
	return getMembershipByUser(cp.init().conn, input)
}

func getMembershipByUserMap(c grpc.ClientConnInterface, input *protos.MemUserGroupInput) (output map[string]*protos.MemberJoin, err error) {
	output = make(map[string]*protos.MemberJoin)
	membership, err := getMembershipByUser(c, input)
	output[membership.ID] = membership
	return output, err
}

// GetMembershipByUserMap ...
func GetMembershipByUserMap(input *protos.MemUserGroupInput) (output map[string]*protos.MemberJoin, err error) {
	return getMembershipByUserMap(use().conn, input)
}

// GetMembershipByUserMap ...
func (cp *ConnProvider) GetMembershipByUserMap(input *protos.MemUserGroupInput) (output map[string]*protos.MemberJoin, err error) {
	return getMembershipByUserMap(cp.init().conn, input)
}
