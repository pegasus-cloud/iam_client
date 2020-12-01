package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func getPermissionByGroup(c grpc.ClientConnInterface, groupID string) (output *protos.PermissionJoinUser, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	permission, err := protos.NewPermissionCRUDControllerClient(c).GetPermissionByGroup(ctx, &protos.PermissionGroupInput{
		GroupID: groupID,
	})
	return permission, err
}

// GetPermissionByGroup ...
func GetPermissionByGroup(groupID string) (output *protos.PermissionJoinUser, err error) {
	return getPermissionByGroup(use().conn, groupID)
}

// GetPermissionByGroup ...
func (cp *ConnProvider) GetPermissionByGroup(groupID string) (output *protos.PermissionJoinUser, err error) {
	return getPermissionByGroup(cp.init().conn, groupID)
}

func getPermissionByGroupMap(c grpc.ClientConnInterface, groupID string) (output map[string]interface{}, err error) {
	permission, err := getPermissionByGroup(c, groupID)
	if err != nil {
		return nil, err
	}
	var permissions []*protos.PermissionJoinUser
	return convert(append(permissions, permission)), nil
}

// GetPermissionByGroupMap ...
func GetPermissionByGroupMap(userID, groupID string) (output map[string]interface{}, err error) {
	return getPermissionByGroupMap(use().conn, groupID)
}

// GetPermissionByGroupMap ...
func (cp *ConnProvider) GetPermissionByGroupMap(userID, groupID string) (output map[string]interface{}, err error) {
	return getPermissionByGroupMap(cp.init().conn, groupID)
}
