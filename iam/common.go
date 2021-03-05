package iam

import (
	"fmt"
	"sync"
	"time"

	"github.com/pegasus-cloud/iam_client/utility"
	"google.golang.org/grpc"
)

const (
	// TCPMode ...
	TCPMode GRPCMode = iota
	// UnixMode ...
	UnixMode
)

type (
	// GRPCMode ...
	GRPCMode int
	// PoolProvider ...
	PoolProvider struct {
		Mode              GRPCMode
		TCPProvider       TCPProvider
		UnixProvider      UnixProvider
		RouteRepsonseType utility.ResponseType
		Timeout           time.Duration
		_                 struct{}
	}
	// TCPProvider ...
	TCPProvider struct {
		Hosts       []string
		ConnPerHost int
		_           struct{}
	}
	// UnixProvider ...
	UnixProvider struct {
		SocketPath string
		ConnCount  int
		_          struct{}
	}
	// ConnProvider ...
	ConnProvider struct {
		Mode              GRPCMode
		Host              string
		SocketPath        string
		RouteRepsonseType utility.ResponseType
		Timeout           time.Duration
		_                 struct{}
	}
	// Pool ...
	pool struct {
		clients chan client
		mu      sync.Mutex
		_       struct{}
	}
	// Client ...
	client struct {
		target  string
		conn    *grpc.ClientConn
		timeout time.Duration
		_       struct{}
	}
)

func (cp *ConnProvider) init() (c client) {
	var err error
	var timeout time.Duration = 5000
	if cp.Timeout != 0 {
		timeout = cp.Timeout
	}
	utility.RouteResponseType = cp.RouteRepsonseType
	switch cp.Mode {
	case UnixMode:
		c.conn, err = grpc.Dial(fmt.Sprintf("unix://%s", cp.SocketPath), grpc.WithInsecure(), grpc.WithTimeout(timeout*time.Millisecond))
	case TCPMode:
		fallthrough
	default:
		c.conn, err = grpc.Dial(cp.Host, grpc.WithInsecure(), grpc.WithTimeout(timeout*time.Millisecond))
	}
	if err != nil {
		panic(err)
	}
	return c
}
