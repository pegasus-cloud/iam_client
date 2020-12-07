package iam

import (
	"net"
	"time"

	"github.com/pegasus-cloud/iam_client/utility"
	"google.golang.org/grpc"
)

var p *pool

// Init ...
func Init(provider PoolProvider) {
	p = &pool{
		hosts:   provider.Hosts,
		count:   len(provider.Hosts) * provider.ConnPerHost,
		clients: make(chan client, len(provider.Hosts)*provider.ConnPerHost),
	}
	var timeout time.Duration = 5000
	if provider.Timeout != 0 {
		timeout = provider.Timeout
	}
	utility.RouteResponseType = provider.RouteRepsonseType
	for i := 0; i < p.count; i++ {
		c, err := grpc.Dial(provider.Hosts[i%len(provider.Hosts)], grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(timeout*time.Millisecond))
		if err != nil {
			panic(err)
		}
		p.clients <- client{
			host:    provider.Hosts[i%len(provider.Hosts)],
			timeout: provider.Timeout,
			conn:    c,
		}
	}
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
	_, err = net.DialTimeout("tcp", c.host, c.timeout*time.Millisecond)
	return
}

func recycle(c client) {
	p.clients <- c
}

// Close ...
func Close() {
	for i := 0; i < p.count; i++ {
		client := <-p.clients
		client.conn.Close()
	}
}
