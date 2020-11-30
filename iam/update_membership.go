package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func updateMembership(c grpc.ClientConnInterface, input *protos.UpdateMembershipInput) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = protos.NewMembershipCURDControllerClient(c).UpdateMembership(ctx, input)
	return err
}

// UpdateMembership ...
func UpdateMembership(input *protos.UpdateMembershipInput) (err error) {
	return updateMembership(use().conn, input)
}

// UpdateMembership ...
func (cp *ConnProvider) UpdateMembership(input *protos.UpdateMembershipInput) (err error) {
	return updateMembership(cp.init().conn, input)
}
