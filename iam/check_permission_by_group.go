package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func checkPermissionByGroup(c grpc.ClientConnInterface, input *protos.PermissionGroupInput) (output *protos.GBoolean, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return protos.NewPermissionCURDControllerClient(c).CheckPermissionByGroup(ctx, input)
}

// CheckPermissionByGroup ...
func CheckPermissionByGroup(input *protos.PermissionGroupInput) (count *protos.GBoolean, err error) {
	return checkPermissionByGroup(use().conn, input)
}

// CheckPermissionByGroup ...
func (cp *ConnProvider) CheckPermissionByGroup(input *protos.PermissionGroupInput) (count *protos.GBoolean, err error) {
	return checkPermissionByGroup(cp.init().conn, input)
}
