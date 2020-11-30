package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func createUser(c grpc.ClientConnInterface, input *protos.UserInfo) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
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

func createUserWithResp(c grpc.ClientConnInterface, input *protos.UserInfo) (output *protos.UserInfo, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return protos.NewUserCURDControllerClient(c).CreateUserWithResp(ctx, input)
}

// CreateUserWithResp ...
func CreateUserWithResp(input *protos.UserInfo) (output *protos.UserInfo, err error) {
	return createUserWithResp(use().conn, input)
}

// CreateUserWithResp ...
func (cp *ConnProvider) CreateUserWithResp(input *protos.UserInfo) (output *protos.UserInfo, err error) {
	return createUserWithResp(cp.init().conn, input)
}

func createUserWithRespMap(c grpc.ClientConnInterface, input *protos.UserInfo) (output map[string]interface{}, err error) {
	user, err := createUserWithResp(c, input)
	var users []*protos.UserInfo
	return convert(append(users, user)), err
}

// CreateUserWithRespMap ...
func CreateUserWithRespMap(input *protos.UserInfo) (output map[string]interface{}, err error) {
	return createUserWithRespMap(use().conn, input)
}

// CreateUserWithRespMap ...
func (cp *ConnProvider) CreateUserWithRespMap(input *protos.UserInfo) (output map[string]interface{}, err error) {
	return createUserWithRespMap(cp.init().conn, input)
}
