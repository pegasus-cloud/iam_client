package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func getGroup(c grpc.ClientConnInterface, input *protos.GroupID) (output *protos.GroupInfo, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return protos.NewGroupCRUDControllerClient(c).GetGroup(ctx, input)
}

// GetGroup ...
func GetGroup(input *protos.GroupID) (output *protos.GroupInfo, err error) {
	return getGroup(use().conn, input)
}

// GetGroup ...
func (cp *ConnProvider) GetGroup(input *protos.GroupID) (output *protos.GroupInfo, err error) {
	return getGroup(cp.init().conn, input)
}

func getGroupMap(c grpc.ClientConnInterface, input *protos.GroupID) (output map[string]interface{}, err error) {
	group, err := getGroup(c, input)
	if err != nil {
		return nil, err
	}
	var groups []*protos.GroupInfo
	return convert(append(groups, group)), nil
}

// GetGroupMap ...
func GetGroupMap(input *protos.GroupID) (output map[string]interface{}, err error) {
	return getGroupMap(use().conn, input)
}

// GetGroupMap ...
func (cp *ConnProvider) GetGroupMap(input *protos.GroupID) (output map[string]interface{}, err error) {
	return getGroupMap(cp.init().conn, input)
}
