import json
import socket


request = {
    "id": 1,
    "params": ["pengze"],
    "method": "rpcServer.Hello"
}

client = socket.create_connection(("localhost", 9000))
client.sendall(json.dumps(request).encode())

# 获取数据
response = client.recv(1024)
response = json.loads(response.decode())
print(response["result"])
