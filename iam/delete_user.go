package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func deleteUser(c grpc.ClientConnInterface, input *protos.UserID) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = protos.NewUserCRUDControllerClient(c).DeleteUser(ctx, input)
	return err
}

// DeleteUser ...
func DeleteUser(input *protos.UserID) (err error) {
	return deleteUser(use().conn, input)
}

// DeleteUser ...
func (cp *ConnProvider) DeleteUser(input *protos.UserID) (err error) {
	return deleteUser(cp.init().conn, input)
}
