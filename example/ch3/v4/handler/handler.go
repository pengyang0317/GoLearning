package handler

// 提升公共命名空间
const RpcServiceName = "handler/rpcServer"

type RpcServer struct{}

func (s *RpcServer) Hello(request string, reply *string) error {
	*reply = "hello, " + request
	return nil
}
