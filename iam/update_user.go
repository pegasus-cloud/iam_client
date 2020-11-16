package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func updateUser(c grpc.ClientConnInterface, input *protos.UpdateInput) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = protos.NewUserCURDControllerClient(c).UpdateUser(ctx, input)
	return err
}

// UpdateUser ...
func UpdateUser(input *protos.UpdateInput) (err error) {
	return updateUser(use().conn, input)
}

// UpdateUser ...
func (cp *ConnProvider) UpdateUser(input *protos.UpdateInput) (err error) {
	return updateUser(cp.init().conn, input)
}
