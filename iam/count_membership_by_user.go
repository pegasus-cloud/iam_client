package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func countMembershipByUser(c grpc.ClientConnInterface, input *protos.UserID) (output *protos.CountOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	output, err = protos.NewMembershipCRUDControllerClient(c).CountMembershipByUser(ctx, input)
	return output, err
}

// CountMembershipByUser ...
func CountMembershipByUser(input *protos.UserID) (count *protos.CountOutput, err error) {
	return countMembershipByUser(use().conn, input)
}

// CountMembershipByUser ...
func (cp *ConnProvider) CountMembershipByUser(input *protos.UserID) (count *protos.CountOutput, err error) {
	return countMembershipByUser(cp.init().conn, input)
}
