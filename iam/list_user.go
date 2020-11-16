package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func listUser(c grpc.ClientConnInterface, limit, offset int) (output *protos.UserInfos, err error) {
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

// ListUser ...
func ListUser(limit, offset int) (output *protos.UserInfos, err error) {
	return listUser(use().conn, limit, offset)
}

// ListUser ...
func (cp *ConnProvider) ListUser(limit, offset int) (output *protos.UserInfos, err error) {
	return listUser(cp.init().conn, limit, offset)
}
