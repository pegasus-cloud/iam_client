package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func updateMembershipByGroup(c grpc.ClientConnInterface, input *protos.UpdateInput) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = protos.NewMembershipCURDControllerClient(c).UpdateMembershipByGroup(ctx, input)
	return err
}

// UpdateMembershipByGroup ...
func UpdateMembershipByGroup(input *protos.UpdateInput) (err error) {
	return updateMembershipByGroup(use().conn, input)
}

// UpdateMembershipByGroup ...
func (cp *ConnProvider) UpdateMembershipByGroup(input *protos.UpdateInput) (err error) {
	return updateMembershipByGroup(cp.init().conn, input)
}
