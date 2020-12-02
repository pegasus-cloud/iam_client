package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func getSecret(c grpc.ClientConnInterface, input *protos.Access) (output *protos.CredentialJoinMembership, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return protos.NewCredentialCRUDControllerClient(c).GetSecret(ctx, input)
}

// GetSecret ...
func GetSecret(input *protos.Access) (output *protos.CredentialJoinMembership, err error) {
	return getSecret(use().conn, input)
}

// GetSecret ...
func (cp *ConnProvider) GetSecret(input *protos.Access) (output *protos.CredentialJoinMembership, err error) {
	return getSecret(cp.init().conn, input)
}

func getSecretMap(c grpc.ClientConnInterface, input *protos.Access) (output map[string]interface{}, err error) {
	credential, err := getSecret(c, input)
	if err != nil {
		return nil, err
	}
	var credentials []*protos.CredentialJoinMembership
	return convert(append(credentials, credential)), nil
}

// GetSecretMap ...
func GetSecretMap(input *protos.Access) (output map[string]interface{}, err error) {
	return getSecretMap(use().conn, input)
}

// GetSecretMap ...
func (cp *ConnProvider) GetSecretMap(input *protos.Access) (output map[string]interface{}, err error) {
	return getSecretMap(cp.init().conn, input)
}
