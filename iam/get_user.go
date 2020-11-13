package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func getUser(c grpc.ClientConnInterface, userID string) (output *protos.UserInfo, err error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	user, err := protos.NewUserCURDControllerClient(use().conn).GetUser(ctx, &protos.GetUserInput{
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

// GetUserMap ...
func GetUserMap(userID string) (output map[string]string, err error) {
	user, err := getUser(use().conn, userID)
	if err != nil {
		return nil, err
	}
	return convertToMap(user), nil
}

// GetUserMap ...
func (cp *ConnProvider) GetUserMap(userID string) (output map[string]string, err error) {
	user, err := getUser(cp.init().conn, userID)
	if err != nil {
		return nil, err
	}
	return convertToMap(user), nil
}
