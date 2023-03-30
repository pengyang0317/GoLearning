import requests


request = { 
    "id": 1,
    "params": ["pengze"],
    "method": "rpcServer.Hello"
}

response = requests.post("http://localhost:9000", json=request)
print(response.json()["result"])
