package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func listUsers(c grpc.ClientConnInterface, input *protos.LimitOffset) (output *protos.ListUserOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return protos.NewUserCRUDControllerClient(c).ListUser(ctx, input)
}

// ListUsers ...
func ListUsers(input *protos.LimitOffset) (output *protos.ListUserOutput, err error) {
	return listUsers(use().conn, input)
}

// ListUsers ...
func (cp *ConnProvider) ListUsers(input *protos.LimitOffset) (output *protos.ListUserOutput, err error) {
	return listUsers(cp.init().conn, input)
}

func listUsersMap(c grpc.ClientConnInterface, input *protos.LimitOffset) (output map[string]*protos.UserInfo, count int64, err error) {
	output = make(map[string]*protos.UserInfo)
	users, err := listUsers(c, input)
	for _, user := range users.Data {
		output[user.ID] = user
	}
	return output, users.Count, err
}

// ListUsersMap ...
func ListUsersMap(input *protos.LimitOffset) (output map[string]*protos.UserInfo, count int64, err error) {
	return listUsersMap(use().conn, input)
}

// ListUsersMap ...
func (cp *ConnProvider) ListUsersMap(input *protos.LimitOffset) (output map[string]*protos.UserInfo, count int64, err error) {
	return listUsersMap(cp.init().conn, input)
}
