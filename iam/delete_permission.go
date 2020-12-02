package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func deletePermission(c grpc.ClientConnInterface, input *protos.PermissionID) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = protos.NewPermissionCRUDControllerClient(c).DeletePermission(ctx, input)
	return err
}

// DeletePermission ...
func DeletePermission(input *protos.PermissionID) (err error) {
	return deletePermission(use().conn, input)
}

// DeletePermission ...
func (cp *ConnProvider) DeletePermission(input *protos.PermissionID) (err error) {
	return deletePermission(cp.init().conn, input)
}
