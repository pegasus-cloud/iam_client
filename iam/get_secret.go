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

func getSecretMap(c grpc.ClientConnInterface, input *protos.Access) (output map[string]*protos.CredentialJoinMembership, err error) {
	credential, err := getSecret(c, input)
	output[credential.UserID] = credential
	return output, err
}

// GetSecretMap ...
func GetSecretMap(input *protos.Access) (output map[string]*protos.CredentialJoinMembership, err error) {
	return getSecretMap(use().conn, input)
}

// GetSecretMap ...
func (cp *ConnProvider) GetSecretMap(input *protos.Access) (output map[string]*protos.CredentialJoinMembership, err error) {
	return getSecretMap(cp.init().conn, input)
}
