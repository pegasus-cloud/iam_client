package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func getUser(c grpc.ClientConnInterface, input *protos.UserID) (output *protos.UserInfo, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return protos.NewUserCRUDControllerClient(c).GetUser(ctx, input)
}

// GetUser ...
func GetUser(input *protos.UserID) (output *protos.UserInfo, err error) {
	return getUser(use().conn, input)
}

// GetUser ...
func (cp *ConnProvider) GetUser(input *protos.UserID) (output *protos.UserInfo, err error) {
	return getUser(cp.init().conn, input)
}

func getUserMap(c grpc.ClientConnInterface, input *protos.UserID) (output map[string]interface{}, err error) {
	user, err := getUser(c, input)
	if err != nil {
		return nil, err
	}
	var users []*protos.UserInfo
	return convert(append(users, user)), nil
}

// GetUserMap ...
func GetUserMap(input *protos.UserID) (output map[string]interface{}, err error) {
	return getUserMap(use().conn, input)
}

// GetUserMap ...
func (cp *ConnProvider) GetUserMap(input *protos.UserID) (output map[string]interface{}, err error) {
	return getUserMap(cp.init().conn, input)
}
