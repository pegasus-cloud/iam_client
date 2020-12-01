package iam

import (
	"fmt"
	"reflect"
	"sync"
	"time"

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

func convert(input interface{}) (output map[string]interface{}) {
	output = make(map[string]interface{})
	switch input.(type) {
	case []*protos.GroupInfo:
		for _, group := range input.([]*protos.GroupInfo) {
			output = toMap(group.ID, reflect.ValueOf(group).Elem())
		}
	case []*protos.UserInfo:
		for _, user := range input.([]*protos.UserInfo) {
			output = toMap(user.ID, reflect.ValueOf(user).Elem())
		}
	case []*protos.MemberJoin:
		for _, membership := range input.([]*protos.MemberJoin) {
			output = toMap(membership.ID, reflect.ValueOf(membership).Elem())
		}
	}
	return output
}

/*
	[Example]
	From:
		GroupInfo struct {
		state         protoimpl.MessageState
		sizeCache     protoimpl.SizeCache
		unknownFields protoimpl.UnknownFields

		ID          string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
		DisplayName string `protobuf:"bytes,2,opt,name=DisplayName,proto3" json:"DisplayName,omitempty"`
	}
	To:
		omit state, sizeCache and unknownFields
		map[<user>.ID] = ID
		map[<user.DisplayName>] = DisplayName
*/
func toMap(prefix string, value reflect.Value) (output map[string]interface{}) {
	output = make(map[string]interface{})
	for i := 3; i < value.NumField(); i++ {
		output[fmt.Sprintf("%s.%s", prefix, value.Type().Field(i).Name)] = value.Field(i).Interface()
	}
	return output
}
