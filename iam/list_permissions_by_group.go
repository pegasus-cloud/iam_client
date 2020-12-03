package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func listPermissionsByGroup(c grpc.ClientConnInterface, input *protos.ListPermissionByGroupInput) (output *protos.ListPermissionJoinOuput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return protos.NewPermissionCRUDControllerClient(c).ListPermissionByGroup(ctx, input)
}

// ListPermissionsByGroup ...
func ListPermissionsByGroup(input *protos.ListPermissionByGroupInput) (output *protos.ListPermissionJoinOuput, err error) {
	return listPermissionsByGroup(use().conn, input)
}

// ListPermissionsByGroup ...
func (cp *ConnProvider) ListPermissionsByGroup(input *protos.ListPermissionByGroupInput) (output *protos.ListPermissionJoinOuput, err error) {
	return listPermissionsByGroup(cp.init().conn, input)
}

func listPermissionsByGroupMap(c grpc.ClientConnInterface, input *protos.ListPermissionByGroupInput) (output map[string]*protos.PermissionJoinUser, count int64, err error) {
	permissions, err := listPermissionsByGroup(c, input)
	for _, permission := range permissions.Data {
		output[permission.ID] = permission
	}
	return output, permissions.Count, err
}

// ListPermissionsByGroupMap ...
func ListPermissionsByGroupMap(input *protos.ListPermissionByGroupInput) (output map[string]*protos.PermissionJoinUser, count int64, err error) {
	return listPermissionsByGroupMap(use().conn, input)
}

// ListPermissionsByGroupMap ...
func (cp *ConnProvider) ListPermissionsByGroupMap(input *protos.ListPermissionByGroupInput) (output map[string]*protos.PermissionJoinUser, count int64, err error) {
	return listPermissionsByGroupMap(cp.init().conn, input)
}
