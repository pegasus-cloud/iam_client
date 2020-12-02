package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func countPermissionByGroup(c grpc.ClientConnInterface, input *protos.GroupID) (output *protos.CountOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	output, err = protos.NewPermissionCRUDControllerClient(c).CountPermissionByGroup(ctx, input)
	return output, err
}

// CountPermissionByGroup ...
func CountPermissionByGroup(input *protos.GroupID) (count *protos.CountOutput, err error) {
	return countPermissionByGroup(use().conn, input)
}

// CountPermissionByGroup ...
func (cp *ConnProvider) CountPermissionByGroup(input *protos.GroupID) (count *protos.CountOutput, err error) {
	return countPermissionByGroup(cp.init().conn, input)
}
