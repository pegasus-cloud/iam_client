package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func getMembership(c grpc.ClientConnInterface, userID, groupID string) (output *protos.MembershipInfo, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	membership, err := protos.NewMembershipCURDControllerClient(use().conn).GetMembership(ctx, &protos.MemUserGroupInput{
		UserID:  userID,
		GroupID: groupID,
	})
	return membership, err
}

// GetMembership ...
func GetMembership(userID, groupID string) (output *protos.MembershipInfo, err error) {
	return getMembership(use().conn, userID, groupID)
}

// GetMembership ...
func (cp *ConnProvider) GetMembership(userID, groupID string) (output *protos.MembershipInfo, err error) {
	return getMembership(cp.init().conn, userID, groupID)
}

func getMembershipMap(c grpc.ClientConnInterface, userID, groupID string) (output map[string]interface{}, err error) {
	membership, err := getMembership(c, userID, groupID)
	if err != nil {
		return nil, err
	}
	var memberships []*protos.MembershipInfo
	return convert(append(memberships, membership)), nil
}

// GetMembershipMap ...
func GetMembershipMap(userID, groupID string) (output map[string]interface{}, err error) {
	return getMembershipMap(use().conn, userID, groupID)
}

// GetMembershipMap ...
func (cp *ConnProvider) GetMembershipMap(userID, groupID string) (output map[string]interface{}, err error) {
	return getMembershipMap(cp.init().conn, userID, groupID)
}
