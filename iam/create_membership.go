package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func createMembership(c grpc.ClientConnInterface, input *protos.MembershipInfo) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = protos.NewMembershipCRUDControllerClient(c).CreateMembership(ctx, input)
	return err
}

// CreateMembership ...
func CreateMembership(input *protos.MembershipInfo) (err error) {
	return createMembership(use().conn, input)
}

// CreateMembership ...
func (cp *ConnProvider) CreateMembership(input *protos.MembershipInfo) (err error) {
	return createMembership(cp.init().conn, input)
}

func createMembershipWithResp(c grpc.ClientConnInterface, input *protos.MembershipInfo) (output *protos.GetMembershipPermissionOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return protos.NewMembershipCRUDControllerClient(c).CreateMembershipWithResp(ctx, input)
}

// CreateMembershipWithResp ...
func CreateMembershipWithResp(input *protos.MembershipInfo) (output *protos.GetMembershipPermissionOutput, err error) {
	return createMembershipWithResp(use().conn, input)
}

// CreateMembershipWithResp ...
func (cp *ConnProvider) CreateMembershipWithResp(input *protos.MembershipInfo) (output *protos.GetMembershipPermissionOutput, err error) {
	return createMembershipWithResp(cp.init().conn, input)
}

func createMembershipWithRespMap(c grpc.ClientConnInterface, input *protos.MembershipInfo) (output map[string]*protos.GetMembershipPermissionOutput, err error) {
	output = make(map[string]*protos.GetMembershipPermissionOutput)
	membership, err := createMembershipWithResp(c, input)
	output[membership.ID] = membership
	return output, err
}

// CreateMembershipWithRespMap ...
func CreateMembershipWithRespMap(input *protos.MembershipInfo) (output map[string]*protos.GetMembershipPermissionOutput, err error) {
	return createMembershipWithRespMap(use().conn, input)
}

// CreateMembershipWithRespMap ...
func (cp *ConnProvider) CreateMembershipWithRespMap(input *protos.MembershipInfo) (output map[string]*protos.GetMembershipPermissionOutput, err error) {
	return createMembershipWithRespMap(cp.init().conn, input)
}
