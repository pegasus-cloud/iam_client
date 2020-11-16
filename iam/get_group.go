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
	return &protos.GroupInfo{
		ID:          group.ID,
		DisplayName: group.DisplayName,
		Description: group.Description,
		Extra:       group.Extra,
		CreatedAt:   group.CreatedAt,
		UpdatedAt:   group.UpdatedAt,
	}, err
}

// GetGroup ...
func GetGroup(groupID string) (output *protos.GroupInfo, err error) {
	return getGroup(use().conn, groupID)
}

// GetGroup ...
func (cp *ConnProvider) GetGroup(groupID string) (output *protos.GroupInfo, err error) {
	return getGroup(cp.init().conn, groupID)
}

func getGroupMap(c grpc.ClientConnInterface, groupID string) (output map[string]string, err error) {
	group, err := getGroup(c, groupID)
	if err != nil {
		return nil, err
	}

	var groups []*protos.GroupInfo
	groups = append(groups, group)
	groupHandler := groupHandler{
		groups: groups,
	}

	return groupHandler.pbToMap(), nil
}

// GetGroupMap ...
func GetGroupMap(groupID string) (output map[string]string, err error) {
	group, err := getGroupMap(use().conn, groupID)
	if err != nil {
		return nil, err
	}
	return group, nil
}

// GetGroupMap ...
func (cp *ConnProvider) GetGroupMap(groupID string) (output map[string]string, err error) {
	group, err := getGroupMap(cp.init().conn, groupID)
	if err != nil {
		return nil, err
	}
	return group, nil
}
