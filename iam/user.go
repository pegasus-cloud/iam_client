package iam

import (
	"context"
	"fmt"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func (cp *ConnProvider) init() (c client) {
	c.conn, _ = grpc.Dial(cp.Host, grpc.WithInsecure(), grpc.WithBlock())
	p.mu.Lock()
	defer p.mu.Unlock()
	return c
}

func createUser(c grpc.ClientConnInterface, input *protos.UserInfo) (err error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	_, err = protos.NewUserCURDControllerClient(c).CreateUser(ctx, input)
	return err
}

// CreateUser ...
func CreateUser(input *protos.UserInfo) (err error) {
	return createUser(use().conn, input)
}

// CreateUser ...
func (cp *ConnProvider) CreateUser(input *protos.UserInfo) (err error) {
	return createUser(cp.init().conn, input)
}

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

func convertToMap(input *protos.UserInfo) (output map[string]string) {
	output = make(map[string]string)
	output[fmt.Sprintf("%s.ID", input.ID)] = input.ID
	output[fmt.Sprintf("%s.DisplayName", input.ID)] = input.DisplayName
	output[fmt.Sprintf("%s.Description", input.ID)] = input.Description
	output[fmt.Sprintf("%s.Extra", input.ID)] = input.Extra
	output[fmt.Sprintf("%s.CreatedAt", input.ID)] = input.CreatedAt
	output[fmt.Sprintf("%s.UpdatedAt", input.ID)] = input.UpdatedAt
	return output
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
