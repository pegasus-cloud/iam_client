package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func createPermission(c grpc.ClientConnInterface, input *protos.PermissionInfo) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = protos.NewPermissionCRUDControllerClient(c).CreatePermission(ctx, input)
	return err
}

// CreatePermission ...
func CreatePermission(input *protos.PermissionInfo) (err error) {
	return createPermission(use().conn, input)
}

// CreatePermission ...
func (cp *ConnProvider) CreatePermission(input *protos.PermissionInfo) (err error) {
	return createPermission(cp.init().conn, input)
}

func createPermissionWithResp(c grpc.ClientConnInterface, input *protos.PermissionInfo) (permission *protos.PermissionJoinUser, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return protos.NewPermissionCRUDControllerClient(c).CreatePermissionWithResp(ctx, input)
}

// CreatePermissionWithResp ...
func CreatePermissionWithResp(input *protos.PermissionInfo) (permission *protos.PermissionJoinUser, err error) {
	return createPermissionWithResp(use().conn, input)
}

// CreatePermissionWithResp ...
func (cp *ConnProvider) CreatePermissionWithResp(input *protos.PermissionInfo) (permission *protos.PermissionJoinUser, err error) {
	return createPermissionWithResp(cp.init().conn, input)
}

func createPermissionWithRespMap(c grpc.ClientConnInterface, input *protos.PermissionInfo) (output map[string]*protos.PermissionJoinUser, err error) {
	permission, err := createPermissionWithResp(c, input)
	output[permission.ID] = permission
	return output, err
}

// CreatePermissionWithRespMap ...
func CreatePermissionWithRespMap(input *protos.PermissionInfo) (output map[string]*protos.PermissionJoinUser, err error) {
	return createPermissionWithRespMap(use().conn, input)
}

// CreatePermissionWithRespMap ...
func (cp *ConnProvider) CreatePermissionWithRespMap(input *protos.PermissionInfo) (output map[string]*protos.PermissionJoinUser, err error) {
	return createPermissionWithRespMap(cp.init().conn, input)
}
