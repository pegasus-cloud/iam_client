package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func getGroup(c grpc.ClientConnInterface, groupID string) (output *protos.GroupInfo, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	group, err := protos.NewGroupCURDControllerClient(use().conn).GetGroup(ctx, &protos.GroupID{
		ID: groupID,
	})
	return group, err
}

// GetGroup ...
func GetGroup(groupID string) (output *protos.GroupInfo, err error) {
	return getGroup(use().conn, groupID)
}

// GetGroup ...
func (cp *ConnProvider) GetGroup(groupID string) (output *protos.GroupInfo, err error) {
	return getGroup(cp.init().conn, groupID)
}

func getGroupMap(c grpc.ClientConnInterface, groupID string) (output map[string]interface{}, err error) {
	group, err := getGroup(c, groupID)
	if err != nil {
		return nil, err
	}
	var groups []*protos.GroupInfo
	return convert(append(groups, group)), nil
}

// GetGroupMap ...
func GetGroupMap(groupID string) (output map[string]interface{}, err error) {
	return getGroupMap(use().conn, groupID)
}

// GetGroupMap ...
func (cp *ConnProvider) GetGroupMap(groupID string) (output map[string]interface{}, err error) {
	return getGroupMap(cp.init().conn, groupID)
}
