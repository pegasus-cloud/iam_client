package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func countMembershipByGroup(c grpc.ClientConnInterface, input *protos.GroupID) (output *protos.CountOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	output, err = protos.NewMembershipCRUDControllerClient(c).CountMembershipByGroup(ctx, input)
	return output, err
}

// CountMembershipByGroup ...
func CountMembershipByGroup(input *protos.GroupID) (count *protos.CountOutput, err error) {
	return countMembershipByGroup(use().conn, input)
}

// CountMembershipByGroup ...
func (cp *ConnProvider) CountMembershipByGroup(input *protos.GroupID) (count *protos.CountOutput, err error) {
	return countMembershipByGroup(cp.init().conn, input)
}
