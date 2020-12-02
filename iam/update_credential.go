package iam

import (
	"context"
	"time"

	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func updateCredential(c grpc.ClientConnInterface, input *protos.UpdateInput) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = protos.NewCredentialCRUDControllerClient(c).UpdateCredential(ctx, input)
	return err
}

// UpdateCredential ...
func UpdateCredential(input *protos.UpdateInput) (err error) {
	return updateCredential(use().conn, input)
}

// UpdateCredential ...
func (cp *ConnProvider) UpdateCredential(input *protos.UpdateInput) (err error) {
	return updateCredential(cp.init().conn, input)
}
