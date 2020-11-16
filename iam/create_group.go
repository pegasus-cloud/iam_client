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
	_, err = protos.NewGroupCURDControllerClient(c).CreateGroup(ctx, input)
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
