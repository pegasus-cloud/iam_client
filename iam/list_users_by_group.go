package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func listUsersByGroup(c grpc.ClientConnInterface, input *protos.GroupID) (output *protos.UserInfos, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return protos.NewUserCRUDControllerClient(c).ListUserByGroup(ctx, input)
}

// ListUsersByGroup ...
func ListUsersByGroup(input *protos.GroupID) (output *protos.UserInfos, err error) {
	return listUsersByGroup(use().conn, input)
}

// ListUsersByGroup ...
func (cp *ConnProvider) ListUsersByGroup(input *protos.GroupID) (output *protos.UserInfos, err error) {
	return listUsersByGroup(cp.init().conn, input)
}

func listUsersByGroupMap(c grpc.ClientConnInterface, input *protos.GroupID) (output map[string]*protos.UserInfo, err error) {
	output = make(map[string]*protos.UserInfo)
	users, err := listUsersByGroup(c, input)
	for _, user := range users.Data {
		output[user.ID] = user
	}
	return output, err
}

// ListUsersByGroupMap ...
func ListUsersByGroupMap(input *protos.GroupID) (output map[string]*protos.UserInfo, err error) {
	return listUsersByGroupMap(use().conn, input)
}

// ListUsersByGroupMap ...
func (cp *ConnProvider) ListUsersByGroupMap(input *protos.GroupID) (output map[string]*protos.UserInfo, err error) {
	return listUsersByGroupMap(cp.init().conn, input)
}
