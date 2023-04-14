import json
import socket


request = {
    "id": 1,
    "params": ["pengze"],
    "method": "rpcServer.Hello"
}
#  连接 RPC 服务端并发送请求数据：使用 socket.create_connection() 方法创建一个 TCP 连接到指定的 RPC 服务端地址和端口号（"localhost", 9000），然后使用 socket.sendall() 方法将请求数据编码为 JSON 格式，并发送给 RPC 服务端。
client = socket.create_connection(("localhost", 9000))
client.sendall(json.dumps(request).encode())

# 接收响应数据：使用 socket.recv() 方法接收 RPC 服务端返回的响应数据，并使用 json.loads() 方法将 JSON 格式的数据解码为 Python 对象。然后从响应数据中获取 "result" 键对应的值，即调用远程方法的返回结果，并使用 print() 方法将其输出。
response = client.recv(1024)
response = json.loads(response.decode())
print(response["result"])
