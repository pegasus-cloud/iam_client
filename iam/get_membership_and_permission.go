package iam

import (
	"context"
	"fmt"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func getMembershipAndPermission(c grpc.ClientConnInterface, userID, groupID string) (output *protos.GetMembershipPermissionOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	fmt.Println(c)
	return protos.NewMembershipCRUDControllerClient(c).GetMembershipPermission(ctx, &protos.MemUserGroupInput{
		UserID:  userID,
		GroupID: groupID,
	})
}

// GetMembershipAndPermission ...
func GetMembershipAndPermission(userID, groupID string) (output *protos.GetMembershipPermissionOutput, err error) {
	return getMembershipAndPermission(use().conn, userID, groupID)
}

// GetMembershipAndPermission ...
func (cp *ConnProvider) GetMembershipAndPermission(userID, groupID string) (output *protos.GetMembershipPermissionOutput, err error) {
	return getMembershipAndPermission(cp.init().conn, userID, groupID)
}

func getMembershipAndPermissionMap(c grpc.ClientConnInterface, userID, groupID string) (output map[string]interface{}, err error) {
	membershipAndPermission, err := getMembershipAndPermission(c, userID, groupID)
	if err != nil {
		return nil, err
	}
	var membershipAndPermissions []*protos.GetMembershipPermissionOutput
	return convert(append(membershipAndPermissions, membershipAndPermission)), nil
}

// GetMembershipAndPermissionMap ...
func GetMembershipAndPermissionMap(userID, groupID string) (output map[string]interface{}, err error) {
	return getMembershipAndPermissionMap(use().conn, userID, groupID)
}

// GetMembershipAndPermissionMap ...
func (cp *ConnProvider) GetMembershipAndPermissionMap(userID, groupID string) (output map[string]interface{}, err error) {
	return getMembershipAndPermissionMap(cp.init().conn, userID, groupID)
}
