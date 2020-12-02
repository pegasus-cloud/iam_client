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

func listGroupsMap(c grpc.ClientConnInterface, input *protos.LimitOffset) (output map[string]interface{}, err error) {
	output = make(map[string]interface{})
	groups, err := listGroups(c, input)
	if err != nil {
		return output, err
	}
	output = convert(groups.Data)
	output["count"] = groups.Count
	return output, nil
}

// ListGroupsMap ...
func ListGroupsMap(input *protos.LimitOffset) (output map[string]interface{}, err error) {
	return listGroupsMap(use().conn, input)
}

// ListGroupsMap ...
func (cp *ConnProvider) ListGroupsMap(input *protos.LimitOffset) (output map[string]interface{}, err error) {
	return listGroupsMap(cp.init().conn, input)
}
