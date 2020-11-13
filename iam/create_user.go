package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func createUser(c grpc.ClientConnInterface, input *protos.UserInfo) (err error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	_, err = protos.NewUserCURDControllerClient(c).CreateUser(ctx, input)
	return err
}

// CreateUser ...
func CreateUser(input *protos.UserInfo) (err error) {
	return createUser(use().conn, input)
}

// CreateUser ...
func (cp *ConnProvider) CreateUser(input *protos.UserInfo) (err error) {
	return createUser(cp.init().conn, input)
}
