package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func updateUser(c grpc.ClientConnInterface, input *protos.UpdateInput) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = protos.NewUserCRUDControllerClient(c).UpdateUser(ctx, input)
	return err
}

// UpdateUser ...
func UpdateUser(input *protos.UpdateInput) (err error) {
	return updateUser(use().conn, input)
}

// UpdateUser ...
func (cp *ConnProvider) UpdateUser(input *protos.UpdateInput) (err error) {
	return updateUser(cp.init().conn, input)
}

func updateUserWithResp(c grpc.ClientConnInterface, input *protos.UpdateInput) (group *protos.UserInfo, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return protos.NewUserCRUDControllerClient(c).UpdateUserWithResp(ctx, input)
}

// UpdateUserWithResp ...
func UpdateUserWithResp(input *protos.UpdateInput) (permission *protos.UserInfo, err error) {
	return updateUserWithResp(use().conn, input)
}

// UpdateUserWithResp ...
func (cp *ConnProvider) UpdateUserWithResp(input *protos.UpdateInput) (permission *protos.UserInfo, err error) {
	return updateUserWithResp(cp.init().conn, input)
}

func updateUserWithRespMap(c grpc.ClientConnInterface, input *protos.UpdateInput) (output map[string]*protos.UserInfo, err error) {
	output = make(map[string]*protos.UserInfo)
	user, err := updateUserWithResp(c, input)
	output[user.ID] = user
	return output, err
}

// UpdateUserWithRespMap ...
func UpdateUserWithRespMap(input *protos.UpdateInput) (output map[string]*protos.UserInfo, err error) {
	return updateUserWithRespMap(use().conn, input)
}

// UpdateUserWithRespMap ...
func (cp *ConnProvider) UpdateUserWithRespMap(input *protos.UpdateInput) (output map[string]*protos.UserInfo, err error) {
	return updateUserWithRespMap(cp.init().conn, input)
}
