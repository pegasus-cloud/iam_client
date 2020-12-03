package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func getMembership(c grpc.ClientConnInterface, input *protos.MemUserGroupInput) (output *protos.MembershipInfo, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return protos.NewMembershipCRUDControllerClient(c).GetMembership(ctx, input)
}

// GetMembership ...
func GetMembership(input *protos.MemUserGroupInput) (output *protos.MembershipInfo, err error) {
	return getMembership(use().conn, input)
}

// GetMembership ...
func (cp *ConnProvider) GetMembership(input *protos.MemUserGroupInput) (output *protos.MembershipInfo, err error) {
	return getMembership(cp.init().conn, input)
}

func getMembershipMap(c grpc.ClientConnInterface, input *protos.MemUserGroupInput) (output map[string]*protos.MembershipInfo, err error) {
	membership, err := getMembership(c, input)
	output[membership.ID] = membership
	return output, err
}

// GetMembershipMap ...
func GetMembershipMap(input *protos.MemUserGroupInput) (output map[string]*protos.MembershipInfo, err error) {
	return getMembershipMap(use().conn, input)
}

// GetMembershipMap ...
func (cp *ConnProvider) GetMembershipMap(input *protos.MemUserGroupInput) (output map[string]*protos.MembershipInfo, err error) {
	return getMembershipMap(cp.init().conn, input)
}
