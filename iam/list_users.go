package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func listUsers(c grpc.ClientConnInterface, limit, offset int) (output *protos.ListUserOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	users, err := protos.NewUserCRUDControllerClient(c).ListUser(ctx, &protos.LimitOffset{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}
	return users, nil
}

// ListUsers ...
func ListUsers(limit, offset int) (output *protos.ListUserOutput, err error) {
	return listUsers(use().conn, limit, offset)
}

// ListUsers ...
func (cp *ConnProvider) ListUsers(limit, offset int) (output *protos.ListUserOutput, err error) {
	return listUsers(cp.init().conn, limit, offset)
}

func listUsersMap(c grpc.ClientConnInterface, limit, offset int) (output map[string]interface{}, err error) {
	output = make(map[string]interface{})
	users, err := listUsers(c, limit, offset)
	if err != nil {
		return output, err
	}
	output = convert(users.Data)
	output["count"] = users.Count
	return output, nil
}

// ListUsersMap ...
func ListUsersMap(limit, offset int) (output map[string]interface{}, err error) {
	groups, err := listUsersMap(use().conn, limit, offset)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

// ListUsersMap ...
func (cp *ConnProvider) ListUsersMap(limit, offset int) (output map[string]interface{}, err error) {
	groups, err := listUsersMap(cp.init().conn, limit, offset)
	if err != nil {
		return nil, err
	}
	return groups, nil
}
