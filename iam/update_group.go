package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func updateGroup(c grpc.ClientConnInterface, input *protos.UpdateInput) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = protos.NewGroupCRUDControllerClient(c).UpdateGroup(ctx, input)
	return err
}

// UpdateGroup ...
func UpdateGroup(input *protos.UpdateInput) (err error) {
	return updateGroup(use().conn, input)
}

// UpdateGroup ...
func (cp *ConnProvider) UpdateGroup(input *protos.UpdateInput) (err error) {
	return updateGroup(cp.init().conn, input)
}

func updateGroupWithResp(c grpc.ClientConnInterface, input *protos.UpdateInput) (group *protos.GroupInfo, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return protos.NewGroupCRUDControllerClient(c).UpdateGroupWithResp(ctx, input)
}

// UpdateGroupWithResp ...
func UpdateGroupWithResp(input *protos.UpdateInput) (permission *protos.GroupInfo, err error) {
	return updateGroupWithResp(use().conn, input)
}

// UpdateGroupWithResp ...
func (cp *ConnProvider) UpdateGroupWithResp(input *protos.UpdateInput) (permission *protos.GroupInfo, err error) {
	return updateGroupWithResp(cp.init().conn, input)
}

func updateGroupWithRespMap(c grpc.ClientConnInterface, input *protos.UpdateInput) (output map[string]*protos.GroupInfo, err error) {
	group, err := updateGroupWithResp(c, input)
	output[group.ID] = group
	return output, err
}

// UpdateGroupWithRespMap ...
func UpdateGroupWithRespMap(input *protos.UpdateInput) (output map[string]*protos.GroupInfo, err error) {
	return updateGroupWithRespMap(use().conn, input)
}

// UpdateGroupWithRespMap ...
func (cp *ConnProvider) UpdateGroupWithRespMap(input *protos.UpdateInput) (output map[string]*protos.GroupInfo, err error) {
	return updateGroupWithRespMap(cp.init().conn, input)
}
