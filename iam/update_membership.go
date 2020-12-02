package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func updateMembership(c grpc.ClientConnInterface, input *protos.UpdateMembershipInput) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = protos.NewMembershipCRUDControllerClient(c).UpdateMembership(ctx, input)
	return err
}

// UpdateMembership ...
func UpdateMembership(input *protos.UpdateMembershipInput) (err error) {
	return updateMembership(use().conn, input)
}

// UpdateMembership ...
func (cp *ConnProvider) UpdateMembership(input *protos.UpdateMembershipInput) (err error) {
	return updateMembership(cp.init().conn, input)
}

func updateMembershipWithResp(c grpc.ClientConnInterface, input *protos.UpdateMembershipInput) (permission *protos.GetMembershipPermissionOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return protos.NewMembershipCRUDControllerClient(c).UpdateMembershipWithResp(ctx, input)
}

// UpdateMembershipWithResp ...
func UpdateMembershipWithResp(input *protos.UpdateMembershipInput) (permission *protos.GetMembershipPermissionOutput, err error) {
	return updateMembershipWithResp(use().conn, input)
}

// UpdateMembershipWithResp ...
func (cp *ConnProvider) UpdateMembershipWithResp(input *protos.UpdateMembershipInput) (permission *protos.GetMembershipPermissionOutput, err error) {
	return updateMembershipWithResp(cp.init().conn, input)
}

func updateMembershipWithRespMap(c grpc.ClientConnInterface, input *protos.UpdateMembershipInput) (output map[string]interface{}, err error) {
	membership, err := updateMembershipWithResp(c, input)
	var memberships []*protos.GetMembershipPermissionOutput
	return convert(append(memberships, membership)), err
}

// UpdateMembershipWithRespMap ...
func UpdateMembershipWithRespMap(input *protos.UpdateMembershipInput) (output map[string]interface{}, err error) {
	return updateMembershipWithRespMap(use().conn, input)
}

// UpdateMembershipWithRespMap ...
func (cp *ConnProvider) UpdateMembershipWithRespMap(input *protos.UpdateMembershipInput) (output map[string]interface{}, err error) {
	return updateMembershipWithRespMap(cp.init().conn, input)
}
