package iam

import (
	"context"
	"fmt"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func listMembershipsByUser(c grpc.ClientConnInterface, userID string, limit, offset int) (output *protos.ListMembershipJoinOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	memberships, err := protos.NewMembershipCRUDControllerClient(c).ListMembershipByUser(ctx, &protos.ListMembershipByUserInput{
		UserID: userID,
		Data: &protos.LimitOffset{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	})
	fmt.Println(limit, offset)
	fmt.Println(memberships)
	if err != nil {
		return nil, err
	}
	return memberships, nil
}

// ListMembershipsByUser ...
func ListMembershipsByUser(userID string, limit, offset int) (output *protos.ListMembershipJoinOutput, err error) {
	return listMembershipsByUser(use().conn, userID, limit, offset)
}

// ListMembershipsByUser ...
func (cp *ConnProvider) ListMembershipsByUser(userID string, limit, offset int) (output *protos.ListMembershipJoinOutput, err error) {
	return listMembershipsByUser(cp.init().conn, userID, limit, offset)
}

func listMembershipsByUserMap(c grpc.ClientConnInterface, userID string, limit, offset int) (output map[string]interface{}, err error) {
	output = make(map[string]interface{})
	memberships, err := listMembershipsByUser(c, userID, limit, offset)
	if err != nil {
		return output, err
	}
	output = convert(memberships.Data)
	output["count"] = memberships.Count
	return output, nil
}

// ListMembershipsByUserMap ...
func ListMembershipsByUserMap(userID string, limit, offset int) (output map[string]interface{}, err error) {
	memberships, err := listMembershipsByUserMap(use().conn, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	return memberships, nil
}

// ListMembershipsByUserMap ...
func (cp *ConnProvider) ListMembershipsByUserMap(userID string, limit, offset int) (output map[string]interface{}, err error) {
	memberships, err := listMembershipsByUserMap(cp.init().conn, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	return memberships, nil
}
