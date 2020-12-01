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
