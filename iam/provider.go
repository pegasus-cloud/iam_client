package iam

import (
	"github.com/pegasus-cloud/iam_client/utility"
	"google.golang.org/grpc"
)

var p *pool

// Init ...
func Init(provider PoolProvider) {
	p = &pool{
		count:   len(provider.Hosts) * provider.ConnPerHost,
		clients: make(chan client, len(provider.Hosts)*provider.ConnPerHost),
	}
	utility.RouteResponseType = provider.RouteRepsonseType
	for i := 0; i < p.count; i++ {
		c, _ := grpc.Dial(provider.Hosts[i%len(provider.Hosts)], grpc.WithInsecure(), grpc.WithBlock())
		p.clients <- client{
			conn: c,
		}
	}
}

// Use ...
func use() (c client) {
	p.mu.Lock()
	c = <-p.clients
	defer recycle(c)
	defer p.mu.Unlock()
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
