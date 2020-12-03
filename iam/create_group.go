package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func createGroup(c grpc.ClientConnInterface, input *protos.GroupInfo) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = protos.NewGroupCRUDControllerClient(c).CreateGroup(ctx, input)
	return err
}

// CreateGroup ...
func CreateGroup(input *protos.GroupInfo) (err error) {
	return createGroup(use().conn, input)
}

// CreateGroup ...
func (cp *ConnProvider) CreateGroup(input *protos.GroupInfo) (err error) {
	return createGroup(cp.init().conn, input)
}

func createGroupWithResp(c grpc.ClientConnInterface, input *protos.GroupInfo) (output *protos.GroupInfo, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return protos.NewGroupCRUDControllerClient(c).CreateGroupWithResp(ctx, input)
}

// CreateGroupWithResp ...
func CreateGroupWithResp(input *protos.GroupInfo) (output *protos.GroupInfo, err error) {
	return createGroupWithResp(use().conn, input)
}

// CreateGroupWithResp ...
func (cp *ConnProvider) CreateGroupWithResp(input *protos.GroupInfo) (output *protos.GroupInfo, err error) {
	return createGroupWithResp(cp.init().conn, input)
}

func createGroupWithRespMap(c grpc.ClientConnInterface, input *protos.GroupInfo) (output map[string]*protos.GroupInfo, err error) {
	group, err := createGroupWithResp(c, input)
	output[group.ID] = group
	return output, err
}

// CreateGroupWithRespMap ...
func CreateGroupWithRespMap(input *protos.GroupInfo) (output map[string]*protos.GroupInfo, err error) {
	return createGroupWithRespMap(use().conn, input)
}

// CreateGroupWithRespMap ...
func (cp *ConnProvider) CreateGroupWithRespMap(input *protos.GroupInfo) (output map[string]*protos.GroupInfo, err error) {
	return createGroupWithRespMap(cp.init().conn, input)
}
