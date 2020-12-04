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

func getCredentialMap(c grpc.ClientConnInterface, input *protos.CredUserGroupInput) (output map[string]*protos.CredentialJoinMembership, err error) {
	output = make(map[string]*protos.CredentialJoinMembership)
	credential, err := getCredential(c, input)
	output[credential.UserID] = credential
	return output, err
}

// GetCredentialMap ...
func GetCredentialMap(input *protos.CredUserGroupInput) (output map[string]*protos.CredentialJoinMembership, err error) {
	return getCredentialMap(use().conn, input)
}

// GetCredentialMap ...
func (cp *ConnProvider) GetCredentialMap(input *protos.CredUserGroupInput) (output map[string]*protos.CredentialJoinMembership, err error) {
	return getCredentialMap(cp.init().conn, input)
}
