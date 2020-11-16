package iam

import (
	"fmt"
	"reflect"
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

type (
	userHandler struct {
		users []*protos.UserInfo
	}
)

func (uh *userHandler) pbToMap() (output map[string]string) {
	output = make(map[string]string)
	for _, user := range uh.users {
		e := reflect.ValueOf(user).Elem()
		for i := 0; i < e.NumField(); i++ {
			output[fmt.Sprintf("%s.%s", user.ID, e.Type().Field(i).Name)] = e.Field(i).Interface().(string)
		}
	}
	return output
}

type (
	groupHandler struct {
		groups []*protos.GroupInfo
	}
)

func (gh *groupHandler) pbToMap() (output map[string]string) {
	output = make(map[string]string)
	for _, group := range gh.groups {
		e := reflect.ValueOf(group).Elem()
		for i := 0; i < e.NumField(); i++ {
			output[fmt.Sprintf("%s.%s", group.ID, e.Type().Field(i).Name)] = e.Field(i).Interface().(string)
		}
	}
	return output
}
