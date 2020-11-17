package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func createMembership(c grpc.ClientConnInterface, input *protos.MembershipInfo) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = protos.NewMembershipCURDControllerClient(c).CreateMembership(ctx, input)
	return err
}

// CreateMembership ...
func CreateMembership(input *protos.MembershipInfo) (err error) {
	return createMembership(use().conn, input)
}

// CreateMembership ...
func (cp *ConnProvider) CreateMembership(input *protos.MembershipInfo) (err error) {
	return createMembership(cp.init().conn, input)
}
