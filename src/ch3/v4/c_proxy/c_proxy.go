package cproxy

import (
	"lgo/src/ch3/v4/handler"
	"net/rpc"
)

type RpcServer struct {
	*rpc.Client
}

func NewRpcServer(addr string) (*RpcServer, error) {
	client, err := rpc.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &RpcServer{Client: client}, nil
}

func (s *RpcServer) Hello(request string, reply *string) error {
	return s.Client.Call(handler.RpcServiceName+".Hello", request, reply)
}
