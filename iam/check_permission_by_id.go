package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func checkPermissionByID(c grpc.ClientConnInterface, input *protos.PermissionID) (output *protos.GBoolean, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return protos.NewPermissionCRUDControllerClient(c).CheckPermissionByID(ctx, input)
}

// CheckPermissionByID ...
func CheckPermissionByID(input *protos.PermissionID) (count *protos.GBoolean, err error) {
	return checkPermissionByID(use().conn, input)
}

// CheckPermissionByID ...
func (cp *ConnProvider) CheckPermissionByID(input *protos.PermissionID) (count *protos.GBoolean, err error) {
	return checkPermissionByID(cp.init().conn, input)
}
