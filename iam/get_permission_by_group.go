package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func getPermissionByGroup(c grpc.ClientConnInterface, input *protos.PermissionGroupInput) (output *protos.PermissionJoinUser, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return protos.NewPermissionCRUDControllerClient(c).GetPermissionByGroup(ctx, input)
}

// GetPermissionByGroup ...
func GetPermissionByGroup(input *protos.PermissionGroupInput) (output *protos.PermissionJoinUser, err error) {
	return getPermissionByGroup(use().conn, input)
}

// GetPermissionByGroup ...
func (cp *ConnProvider) GetPermissionByGroup(input *protos.PermissionGroupInput) (output *protos.PermissionJoinUser, err error) {
	return getPermissionByGroup(cp.init().conn, input)
}

func getPermissionByGroupMap(c grpc.ClientConnInterface, input *protos.PermissionGroupInput) (output map[string]*protos.PermissionJoinUser, err error) {
	permission, err := getPermissionByGroup(c, input)
	output[permission.ID] = permission
	return output, err
}

// GetPermissionByGroupMap ...
func GetPermissionByGroupMap(input *protos.PermissionGroupInput) (output map[string]*protos.PermissionJoinUser, err error) {
	return getPermissionByGroupMap(use().conn, input)
}

// GetPermissionByGroupMap ...
func (cp *ConnProvider) GetPermissionByGroupMap(input *protos.PermissionGroupInput) (output map[string]*protos.PermissionJoinUser, err error) {
	return getPermissionByGroupMap(cp.init().conn, input)
}
