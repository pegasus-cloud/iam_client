package iam

import (
	"fmt"
	"sync"

	"github.com/pegasus-cloud/iam_client/protos"
	"github.com/pegasus-cloud/iam_client/utility"
	"google.golang.org/grpc"
)

type (
	// PoolProvider ...
	PoolProvider struct {
		Hosts             []string
		ConnPerHost       int
		RouteRepsonseType utility.ResponseType
		_                 struct{}
	}
	// ConnProvider ...
	ConnProvider struct {
		Host string
		_    struct{}
	}
	// Pool ...
	pool struct {
		clients chan client
		count   int
		mu      sync.Mutex
		_       struct{}
	}
	// Client ...
	client struct {
		conn *grpc.ClientConn
		_    struct{}
	}
)

func (cp *ConnProvider) init() (c client) {
	c.conn, _ = grpc.Dial(cp.Host, grpc.WithInsecure(), grpc.WithBlock())
	p.mu.Lock()
	defer p.mu.Unlock()
	return c
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
