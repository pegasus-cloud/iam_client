package iam

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func countUser(c grpc.ClientConnInterface) (output *protos.CountOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	output, err = protos.NewUserCURDControllerClient(c).CountUser(ctx, &empty.Empty{})
	return output, err
}

// CountUser ...
func CountUser() (count *protos.CountOutput, err error) {
	return countUser(use().conn)
}

// CountUser ...
func (cp *ConnProvider) CountUser() (count *protos.CountOutput, err error) {
	return countUser(cp.init().conn)
}
