package iam

import (
	"sync"

	"google.golang.org/grpc"
)

type (
	// PoolProvider ...
	PoolProvider struct {
		Hosts       []string
		ConnPerHost int
		_           struct{}
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
