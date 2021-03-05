package iam

import (
	"fmt"
	"time"

	"github.com/pegasus-cloud/iam_client/utility"
	"google.golang.org/grpc"
)

var p *pool

// Init ...
func Init(provider PoolProvider) {
	p = &pool{}
	var timeout time.Duration = 5000
	if provider.Timeout != 0 {
		timeout = provider.Timeout
	}
	utility.RouteResponseType = provider.RouteRepsonseType
	targets := []string{}

	switch provider.Mode {
	case UnixMode:
		count := provider.UnixProvider.ConnCount
		p.clients = make(chan client, count)
		connTerm := fmt.Sprintf("unix://%s", provider.UnixProvider.SocketPath)
		for i := 0; i < count; i++ {
			targets = append(targets, connTerm)
		}
		if err := dial(UnixMode, targets, timeout); err != nil {
			panic(err)
		}
	case TCPMode:
		fallthrough
	default:
		count := len(provider.TCPProvider.Hosts) * provider.TCPProvider.ConnPerHost
		p.clients = make(chan client, count)
		for i := 0; i < count; i++ {
			targets = append(targets, provider.TCPProvider.Hosts[i%len(provider.TCPProvider.Hosts)])
		}
		if err := dial(TCPMode, targets, timeout); err != nil {
			panic(err)
		}
	}
}

func dial(mode GRPCMode, targets []string, timeout time.Duration) (err error) {
	for _, target := range targets {
		c, err := grpc.Dial(target, grpc.WithInsecure(), grpc.WithTimeout(timeout*time.Millisecond))
		if err != nil {
			return err
		}
		p.clients <- client{
			target:  target,
			timeout: timeout,
			conn:    c,
		}
	}
	return nil
}

// Use ...
func use() (c client) {
	p.mu.Lock()
	for {
		c = get()
		if check(c) == nil {
			break
		}
		recycle(c)
	}
	defer recycle(c)
	defer p.mu.Unlock()
	return
}

func get() (c client) {
	c = <-p.clients
	return
}

func check(c client) (err error) {
	_, err = grpc.Dial(c.target, grpc.WithInsecure(), grpc.WithTimeout(c.timeout*time.Millisecond))
	return
}

func recycle(c client) {
	p.clients <- c
}

// Close ...
func Close() {
	for i := 0; i < len(p.clients); i++ {
		client := <-p.clients
		client.conn.Close()
		p = nil
	}
}
