package sproxy

import (
	"lgo/src/ch3/v4/handler"
	"net/rpc"
)



type rpcServerI interface {
	Hello(request string, reply *string) error
}


func RegisterRpcServer(s rpcServerI) {
	rpc.RegisterName(handler.RpcServiceName, s)
}
