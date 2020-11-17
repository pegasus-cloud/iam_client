package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func listUsers(c grpc.ClientConnInterface, limit, offset int) (output *protos.UserInfos, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	users, err := protos.NewUserCURDControllerClient(c).ListUser(ctx, &protos.LimitOffset{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}
	return users, nil
}

// ListUsers ...
func ListUsers(limit, offset int) (output *protos.UserInfos, err error) {
	return listUsers(use().conn, limit, offset)
}

// ListUsers ...
func (cp *ConnProvider) ListUsers(limit, offset int) (output *protos.UserInfos, err error) {
	return listUsers(cp.init().conn, limit, offset)
}
