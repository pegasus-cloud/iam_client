package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func deleteMembership(c grpc.ClientConnInterface, input *protos.MemUserGroupInput) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = protos.NewMembershipCURDControllerClient(c).DeleteMembership(ctx, input)
	return err
}

// DeleteMembership ...
func DeleteMembership(input *protos.MemUserGroupInput) (err error) {
	return deleteMembership(use().conn, input)
}

// DeleteMembership ...
func (cp *ConnProvider) DeleteMembership(input *protos.MemUserGroupInput) (err error) {
	return deleteMembership(cp.init().conn, input)
}
