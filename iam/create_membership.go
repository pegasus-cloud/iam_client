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
	_, err = protos.NewMembershipCURDControllerClient(c).CreateMembership(ctx, input)
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
	return protos.NewMembershipCURDControllerClient(c).CreateMembershipWithResp(ctx, input)
}

// CreateMembershipWithResp ...
func CreateMembershipWithResp(input *protos.MembershipInfo) (output *protos.GetMembershipPermissionOutput, err error) {
	return createMembershipWithResp(use().conn, input)
}

// CreateMembershipWithResp ...
func (cp *ConnProvider) CreateMembershipWithResp(input *protos.MembershipInfo) (output *protos.GetMembershipPermissionOutput, err error) {
	return createMembershipWithResp(cp.init().conn, input)
}

func createMembershipWithRespMap(c grpc.ClientConnInterface, input *protos.MembershipInfo) (output map[string]interface{}, err error) {
	membership, err := createMembershipWithResp(c, input)
	var memberships []*protos.GetMembershipPermissionOutput
	return convert(append(memberships, membership)), err
}

// CreateMembershipWithRespMap ...
func CreateMembershipWithRespMap(input *protos.MembershipInfo) (output map[string]interface{}, err error) {
	return createMembershipWithRespMap(use().conn, input)
}

// CreateMembershipWithRespMap ...
func (cp *ConnProvider) CreateMembershipWithRespMap(input *protos.MembershipInfo) (output map[string]interface{}, err error) {
	return createMembershipWithRespMap(cp.init().conn, input)
}
