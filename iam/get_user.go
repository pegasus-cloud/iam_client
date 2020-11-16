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
	user, err := protos.NewUserCURDControllerClient(use().conn).GetUser(ctx, &protos.UserID{
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

func getUserMap(c grpc.ClientConnInterface, userID string) (output map[string]string, err error) {
	user, err := getUser(c, userID)
	if err != nil {
		return nil, err
	}

	var users []*protos.UserInfo
	users = append(users, user)
	userHandler := userHandler{
		users: users,
	}

	return userHandler.pbToMap(), nil
}

// GetUserMap ...
func GetUserMap(userID string) (output map[string]string, err error) {
	user, err := getUserMap(use().conn, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserMap ...
func (cp *ConnProvider) GetUserMap(userID string) (output map[string]string, err error) {
	user, err := getUserMap(cp.init().conn, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
