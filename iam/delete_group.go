package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func deleteGroup(c grpc.ClientConnInterface, input *protos.GroupID) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = protos.NewGroupCRUDControllerClient(c).DeleteGroup(ctx, input)
	return err
}

// DeleteGroup ...
func DeleteGroup(input *protos.GroupID) (err error) {
	return deleteGroup(use().conn, input)
}

// DeleteGroup ...
func (cp *ConnProvider) DeleteGroup(input *protos.GroupID) (err error) {
	return deleteGroup(cp.init().conn, input)
}
