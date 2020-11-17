package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func listMembershipsByUser(c grpc.ClientConnInterface, userID string, limit, offset int) (output *protos.ListMembershipByUserOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	memberships, err := protos.NewMembershipCURDControllerClient(c).ListMembershipByUser(ctx, &protos.ListMembershipByUserInput{
		UserID: userID,
		Data: &protos.LimitOffset{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	})
	if err != nil {
		return nil, err
	}
	return memberships, nil
}

// ListMembershipsByUser ...
func ListMembershipsByUser(userID string, limit, offset int) (output *protos.ListMembershipByUserOutput, err error) {
	return listMembershipsByUser(use().conn, userID, limit, offset)
}

// ListMembershipsByUser ...
func (cp *ConnProvider) ListMembershipsByUser(userID string, limit, offset int) (output *protos.ListMembershipByUserOutput, err error) {
	return listMembershipsByUser(cp.init().conn, userID, limit, offset)
}
