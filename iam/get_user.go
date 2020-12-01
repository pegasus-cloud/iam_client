package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func getUser(c grpc.ClientConnInterface, userID string) (output *protos.UserInfo, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	user, err := protos.NewUserCRUDControllerClient(c).GetUser(ctx, &protos.UserID{
		ID: userID,
	})
	return &protos.UserInfo{
		ID:          user.ID,
		DisplayName: user.DisplayName,
		Description: user.Description,
		Extra:       user.Extra,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}, err
}

// GetUser ...
func GetUser(userID string) (output *protos.UserInfo, err error) {
	return getUser(use().conn, userID)
}

// GetUser ...
func (cp *ConnProvider) GetUser(userID string) (output *protos.UserInfo, err error) {
	return getUser(cp.init().conn, userID)
}

func getUserMap(c grpc.ClientConnInterface, userID string) (output map[string]interface{}, err error) {
	user, err := getUser(c, userID)
	if err != nil {
		return nil, err
	}
	var users []*protos.UserInfo
	return convert(append(users, user)), nil
}

// GetUserMap ...
func GetUserMap(userID string) (output map[string]interface{}, err error) {
	return getUserMap(use().conn, userID)
}

// GetUserMap ...
func (cp *ConnProvider) GetUserMap(userID string) (output map[string]interface{}, err error) {
	return getUserMap(cp.init().conn, userID)
}
