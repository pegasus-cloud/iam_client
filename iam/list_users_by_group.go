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

func listUsersByGroupMap(c grpc.ClientConnInterface, input *protos.GroupID) (output map[string]interface{}, err error) {
	output = make(map[string]interface{})
	users, err := listUsersByGroup(c, input)
	if err != nil {
		return output, err
	}
	output = convert(users.Data)
	return output, nil
}

// ListUsersByGroupMap ...
func ListUsersByGroupMap(input *protos.GroupID) (output map[string]interface{}, err error) {
	return listUsersByGroupMap(use().conn, input)
}

// ListUsersByGroupMap ...
func (cp *ConnProvider) ListUsersByGroupMap(input *protos.GroupID) (output map[string]interface{}, err error) {
	return listUsersByGroupMap(cp.init().conn, input)
}
