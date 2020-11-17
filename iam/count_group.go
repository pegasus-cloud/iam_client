package iam

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pegasus-cloud/iam_client/protos"
	"google.golang.org/grpc"
)

func countGroup(c grpc.ClientConnInterface) (output *protos.CountOutput, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	output, err = protos.NewGroupCURDControllerClient(c).CountGroup(ctx, &empty.Empty{})
	return output, err
}

// CountGroup ...
func CountGroup() (count *protos.CountOutput, err error) {
	return countGroup(use().conn)
}

// CountGroup ...
func (cp *ConnProvider) CountGroup() (count *protos.CountOutput, err error) {
	return countGroup(cp.init().conn)
}
