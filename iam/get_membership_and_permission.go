package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func getMembershipAndPermission(c grpc.ClientConnInterface, input *protos.MemUserGroupInput) (output *protos.GetMembershipPermissionOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return protos.NewMembershipCRUDControllerClient(c).GetMembershipPermission(ctx, input)
}

// GetMembershipAndPermission ...
func GetMembershipAndPermission(input *protos.MemUserGroupInput) (output *protos.GetMembershipPermissionOutput, err error) {
	return getMembershipAndPermission(use().conn, input)
}

// GetMembershipAndPermission ...
func (cp *ConnProvider) GetMembershipAndPermission(input *protos.MemUserGroupInput) (output *protos.GetMembershipPermissionOutput, err error) {
	return getMembershipAndPermission(cp.init().conn, input)
}

func getMembershipAndPermissionMap(c grpc.ClientConnInterface, input *protos.MemUserGroupInput) (output map[string]*protos.GetMembershipPermissionOutput, err error) {
	membershipAndPermission, err := getMembershipAndPermission(c, input)
	output[membershipAndPermission.ID] = membershipAndPermission
	return output, err
}

// GetMembershipAndPermissionMap ...
func GetMembershipAndPermissionMap(input *protos.MemUserGroupInput) (output map[string]*protos.GetMembershipPermissionOutput, err error) {
	return getMembershipAndPermissionMap(use().conn, input)
}

// GetMembershipAndPermissionMap ...
func (cp *ConnProvider) GetMembershipAndPermissionMap(input *protos.MemUserGroupInput) (output map[string]*protos.GetMembershipPermissionOutput, err error) {
	return getMembershipAndPermissionMap(cp.init().conn, input)
}
