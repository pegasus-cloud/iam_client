package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func getFrozen(c grpc.ClientConnInterface, input *protos.MemUserGroupInput) (output *protos.GBoolean, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return protos.NewMembershipCRUDControllerClient(c).GetFrozen(ctx, input)
}

// GetFrozen ...
func GetFrozen(input *protos.MemUserGroupInput) (output *protos.GBoolean, err error) {
	return getFrozen(use().conn, input)
}

// GetFrozen ...
func (cp *ConnProvider) GetFrozen(input *protos.MemUserGroupInput) (output *protos.GBoolean, err error) {
	return getFrozen(cp.init().conn, input)
}
