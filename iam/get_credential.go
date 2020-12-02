package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func getCredential(c grpc.ClientConnInterface, input *protos.CredUserGroupInput) (output *protos.CredentialJoinMembership, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return protos.NewCredentialCRUDControllerClient(c).GetCredential(ctx, input)
}

// GetCredential ...
func GetCredential(input *protos.CredUserGroupInput) (output *protos.CredentialJoinMembership, err error) {
	return getCredential(use().conn, input)
}

// GetCredential ...
func (cp *ConnProvider) GetCredential(input *protos.CredUserGroupInput) (output *protos.CredentialJoinMembership, err error) {
	return getCredential(cp.init().conn, input)
}

func getCredentialMap(c grpc.ClientConnInterface, input *protos.CredUserGroupInput) (output map[string]interface{}, err error) {
	credential, err := getCredential(c, input)
	if err != nil {
		return nil, err
	}
	var credentials []*protos.CredentialJoinMembership
	return convert(append(credentials, credential)), nil
}

// GetCredentialMap ...
func GetCredentialMap(input *protos.CredUserGroupInput) (output map[string]interface{}, err error) {
	return getCredentialMap(use().conn, input)
}

// GetCredentialMap ...
func (cp *ConnProvider) GetCredentialMap(input *protos.CredUserGroupInput) (output map[string]interface{}, err error) {
	return getCredentialMap(cp.init().conn, input)
}
