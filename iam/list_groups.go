package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func listGroups(c grpc.ClientConnInterface, input *protos.LimitOffset) (output *protos.ListGroupOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return protos.NewGroupCRUDControllerClient(c).ListGroup(ctx, input)
}

// ListGroups ...
func ListGroups(input *protos.LimitOffset) (output *protos.ListGroupOutput, err error) {
	return listGroups(use().conn, input)
}

// ListGroups ...
func (cp *ConnProvider) ListGroups(input *protos.LimitOffset) (output *protos.ListGroupOutput, err error) {
	return listGroups(cp.init().conn, input)
}

func listGroupsMap(c grpc.ClientConnInterface, input *protos.LimitOffset) (output map[string]*protos.GroupInfo, count int64, err error) {
	output = make(map[string]*protos.GroupInfo)
	groups, err := listGroups(c, input)
	for _, group := range groups.Data {
		output[group.ID] = group
	}
	return output, groups.Count, err
}

// ListGroupsMap ...
func ListGroupsMap(input *protos.LimitOffset) (output map[string]*protos.GroupInfo, count int64, err error) {
	return listGroupsMap(use().conn, input)
}

// ListGroupsMap ...
func (cp *ConnProvider) ListGroupsMap(input *protos.LimitOffset) (output map[string]*protos.GroupInfo, count int64, err error) {
	return listGroupsMap(cp.init().conn, input)
}
