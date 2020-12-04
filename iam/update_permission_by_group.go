package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func updatePermissionByGroup(c grpc.ClientConnInterface, input *protos.UpdatePermissionByGroupInput) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = protos.NewPermissionCRUDControllerClient(c).UpdatePermissionByGroup(ctx, input)
	return err
}

// UpdatePermissionByGroup ...
func UpdatePermissionByGroup(input *protos.UpdatePermissionByGroupInput) (err error) {
	return updatePermissionByGroup(use().conn, input)
}

// UpdatePermissionByGroup ...
func (cp *ConnProvider) UpdatePermissionByGroup(input *protos.UpdatePermissionByGroupInput) (err error) {
	return updatePermissionByGroup(cp.init().conn, input)
}

func updatePermissionByGroupWithResp(c grpc.ClientConnInterface, input *protos.UpdatePermissionByGroupInput) (permission *protos.PermissionJoinUser, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return protos.NewPermissionCRUDControllerClient(c).UpdatePermissionByGroupWithResp(ctx, input)
}

// UpdatePermissionByGroupWithResp ...
func UpdatePermissionByGroupWithResp(input *protos.UpdatePermissionByGroupInput) (permission *protos.PermissionJoinUser, err error) {
	return updatePermissionByGroupWithResp(use().conn, input)
}

// UpdatePermissionByGroupWithResp ...
func (cp *ConnProvider) UpdatePermissionByGroupWithResp(input *protos.UpdatePermissionByGroupInput) (permission *protos.PermissionJoinUser, err error) {
	return updatePermissionByGroupWithResp(cp.init().conn, input)
}

func updatePermissionByGroupWithRespMap(c grpc.ClientConnInterface, input *protos.UpdatePermissionByGroupInput) (output map[string]*protos.PermissionJoinUser, err error) {
	output = make(map[string]*protos.PermissionJoinUser)
	permission, err := updatePermissionByGroupWithResp(c, input)
	output[permission.ID] = permission
	return output, err
}

// UpdatePermissionByGroupWithRespMap ...
func UpdatePermissionByGroupWithRespMap(input *protos.UpdatePermissionByGroupInput) (output map[string]*protos.PermissionJoinUser, err error) {
	return updatePermissionByGroupWithRespMap(use().conn, input)
}

// UpdatePermissionByGroupWithRespMap ...
func (cp *ConnProvider) UpdatePermissionByGroupWithRespMap(input *protos.UpdatePermissionByGroupInput) (output map[string]*protos.PermissionJoinUser, err error) {
	return updatePermissionByGroupWithRespMap(cp.init().conn, input)
}
