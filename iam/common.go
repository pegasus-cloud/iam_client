package iam

import (
	"sync"
	"time"

	"github.com/pegasus-cloud/iam_client/utility"
	"google.golang.org/grpc"
)

type (
	// PoolProvider ...
	PoolProvider struct {
		Hosts             []string
		ConnPerHost       int
		RouteRepsonseType utility.ResponseType
		Timeout           time.Duration
		_                 struct{}
	}
	// ConnProvider ...
	ConnProvider struct {
		Host    string
		Timeout time.Duration
		_       struct{}
	}
	// Pool ...
	pool struct {
		hosts   []string
		clients chan client
		count   int
		mu      sync.Mutex
		_       struct{}
	}
	// Client ...
	client struct {
		host    string
		conn    *grpc.ClientConn
		timeout time.Duration
		_       struct{}
	}
)

func (cp *ConnProvider) init() (c client) {
	var err error
	c.conn, err = grpc.Dial(cp.Host, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(cp.Timeout*time.Millisecond))
	if err != nil {
		panic(err)
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	return c
}
