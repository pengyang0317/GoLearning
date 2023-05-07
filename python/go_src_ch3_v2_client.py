import json
import socket

# 发送数据

request = {
    "id": 1,
    "params": ["pengze"],
    "method": "rpcServer.Hello"
}
# 创建socket连接
client = socket.create_connection(("localhost", 9000))
client.sendall(json.dumps(request).encode())

# 获取数据
response = client.recv(1024)
response = json.loads(response.decode())
print(response["result"])
