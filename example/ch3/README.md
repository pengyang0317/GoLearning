
# go内置的RPC快速开发

[TOC]

完整代码地址： <https://github1s.com/pengyang0317/GoLearning/blob/main/src/ch3/README.md>

### 1. 什么是RPC

RPC（Remote Procedure Call，远程过程调用）是一种计算机通信协议，用于在不同的进程或者不同的机器之间进行数据交换和调用远程服务。RPC 允许一个程序在一台计算机上调用另一个程序或者子程序，而调用方无需了解被调用方的具体实现细节，只需要知道被调用方的服务名和方法名即可。

在 RPC 中，通常会有一个客户端和一个服务端，客户端通过网络向服务端发送请求，服务端接收请求并处理，然后将结果返回给客户端。RPC 的实现方式有很多种，包括基于 HTTP、TCP、UDP 等不同的协议，也包括基于不同编程语言和框架的实现。

RPC 的优点是可以将分布式系统中的不同部分组合成一个整体，提高了系统的可扩展性和灵活性，并且可以隐藏底层的实现细节，简化了编程的复杂度。RPC 在分布式系统、微服务架构、大规模数据处理等领域中得到了广泛的应用。

### 2.远程过程调用带来的新问题

**2.1Call ID映射**

我们怎么告诉远程机器我们要调用add，而不是del或者FooBar呢？

在本地调用中，函数体是直接通过函数指针来指定的，我们调用add，编译器就自动帮我们调用它相应的函数指针。但是在远程调用中，函数指针是不行的，因为两个进程的地址空间是完全不一样的。

所以，在RPC中，所有的函数都必须有自己的一个ID。这个ID在所有进程中都是唯一确定的。

客户端在做远程过程调用时，必须附上这个ID。然后我们还需要在客户端和服务端分别维护一个 {函数 <--> Call ID} 的对应表。两者的表不一定需要完全相同，但相同的函数对应的Call ID必须相同。

当客户端需要进行远程调用时，它就查一下这个表，找出相应的Call ID，然后把它传给服务端，服务端也通过查表，来确定客户端需要调用的函数，然后执行相应函数的代码。

**2.2序列化和反序列化**
客户端怎么把参数值传给远程的函数呢？

在本地调用中，我们只需要把参数压到栈里，然后让函数自己去栈里读就行。

但是在远程过程调用时，客户端跟服务端是不同的进程，不能通过内存来传递参数。甚至有时候客户端和服务端使用的都不是同一种语言。

这时候就需要客户端把参数先转成一个字节流，传给服务端后，再把字节流转成自己能读取的格式。这个过程叫序列化和反序列化。

同理，从服务端返回的值也需要序列化反序列化的过程。

**2.3网络传输**
远程调用往往用在网络上，客户端和服务端是通过网络连接的。所有的数据都需要通过网络传输，因此就需要有一个网络传输层。

网络传输层需要把Call ID 和 序列化后的参数字节流传给服务端，然后再把序列化后的调用结果传回客户端。

只要能完成这两者的，都可以作为传输层使用。因此，它所使用的协议其实是不限的，能完成传输就行。尽管大部分RPC框架都使用TCP协议，但其实UDP也可以，而gRPC干脆就用了HTTP2。Java的Netty也属于这层的东西。

有了这三个机制，就能实现RPC了

### 3.基于 TCP 协议实现简单的RPC

 这段 Go 代码实现了一个基于 TCP 协议的 RPC 服务端 代码位置： ch3/v1/server

```go
const PORT = ":9000"

type rpcServer struct{}

func (s *rpcServer) Hello(request string, reply *string) error {
 *reply = "hello, " + request
 return nil
}

func main() {
 // 1.实例化一个 server：使用 net.Listen 方法监听指定端口（PORT），创建一个 TCP 的 listener 对象，用于接收客户端的连接请求。
 listener, _ := net.Listen("tcp", PORT)
 // 2.注册处理逻辑：使用 rpc.RegisterName 方法注册一个名为 "rpcServer" 的服务，并将其关联到一个实现了对应方法的结构体 rpcServer 上。
 rpc.RegisterName("rpcServer", &rpcServer{})

 //3.启动服务：调用 listener.Accept() 方法接受客户端的连接请求，并通过 rpc.ServeConn 方法将客户端的请求交给 RPC 服务端处理。
 conn, _ := listener.Accept()
 rpc.ServeConn(conn)
}

```

这段 Go 代码实现了一个基于 TCP 协议的 RPC 客户端， 代码位置 ch3/v1/client

```go

const PORT = ":9000"

func main() {
 // 1.连接 RPC 服务端：使用 rpc.Dial 方法连接指定的 RPC 服务端，传入 tcp 协议和服务端的端口号（PORT）。
 client, err := rpc.Dial("tcp", PORT)
 if err != nil {
  fmt.Printf("dial error: %v", err)
 }
 var reply string

 // 2.调用远程方法：使用 client.Call 方法调用远程方法，传入方法名、参数和返回值指针。这里的方法名为 "rpcServer.Hello"，表示调用名为 "Hello" 的方法，该方法属于名为 "rpcServer" 的 RPC 服务。参数为 "pengze"，表示将 "pengze" 作为参数传递给远程方法。返回值为 reply，是一个字符串类型的指针，用于接收远程方法的返回结果。
 err = client.Call("rpcServer.Hello", "pengze", &reply)

 if err != nil {
  fmt.Printf("call error: %v", err)
 }
 fmt.Println(reply)

}
```

### 4.基于 TCP 协议和 JSON-RPC 编码的 RPC 服务端 python、java客户端

这段 Go 代码实现了一个基于 TCP 协议和 JSON-RPC 编码的 RPC 服务端  代码位置 ch3/v2/server

```go

type rpcServer struct{}

func (s *rpcServer) Hello(request string, reply *string) error {
 *reply = "hello, " + request
 return nil
}

const PORT = ":9000"

func main() {
 // 1.实例化一个 server：使用 net.Listen 方法监听指定端口 9000，创建一个 TCP 的 listener 对象，用于接收客户端的连接请求。
 listener, _ := net.Listen("tcp", PORT)
 // 2.注册处理逻辑：使用 rpc.RegisterName 方法注册一个名为 "rpcServer" 的服务，并将其关联到一个实现了对应方法的结构体 rpcServer 上
 rpc.RegisterName("rpcServer", &rpcServer{})

 //3.启动服务：使用 listener.Accept() 方法阻塞等待客户端的连接请求，然后通过 jsonrpc.NewServerCodec 方法创建一个 JSON-RPC 编码的服务器编解码器，并通过 rpc.ServeCodec 方法将客户端的请求交给 RPC 服务端处理。由于 listener.Accept() 方法是阻塞的，需要使用一个 for 循环不断接收客户端的连接请求。
 for {
  conn, _ := listener.Accept()
  go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
 }

}

```

这段 Python 代码实现了一个基于 Socket 和 JSON-RPC 编码的 RPC 客户端

``` python
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
```

以下使用 Java 14 实现 RPC 客户端：

maven

```java
<dependency>
    <groupId>com.google.code.gson</groupId>
    <artifactId>gson</artifactId>
    <version>2.8.9</version>
</dependency>

```

java代码

```java
package com.pengze;

import com.google.gson.Gson;

import java.io.*;
import java.net.*;
import java.util.*;

public class RpcClient {

    public static void main(String[] args) throws Exception {
        // 创建请求数据
        Map<String, Object> request = new HashMap<>();
        request.put("id", 1);
        request.put("params", new String[] {"pengze"});
        request.put("method", "rpcServer.Hello");

        // 创建 TCP 连接并发送请求数据
        Socket socket = new Socket("localhost", 9000);
        OutputStream outputStream = socket.getOutputStream();
        PrintWriter writer = new PrintWriter(outputStream);
        writer.println(new Gson().toJson(request));
        writer.flush();

        // 接收响应数据并输出结果
        InputStream inputStream = socket.getInputStream();
        BufferedReader reader = new BufferedReader(new InputStreamReader(inputStream));
        StringBuilder responseBuilder = new StringBuilder();
        responseBuilder.append(reader.readLine());
        Map<String, Object> response = new Gson().fromJson(responseBuilder.toString(), Map.class);
        Object result = response.get("result");
        System.out.println(result);

        // 关闭连接
        socket.close();
    }
}

```

java代码不放到git上了，这是我的输出结果
![352661cb8a5b5a202476509474a4cdc4.png](evernotecid://2DF34E2B-D627-4C9D-B047-ED2B8738D0D2/appyinxiangcom/25024422/ENResource/p1629)

### 5.基于http协议完成add服务端功能

这段 Go 代码实现了一个基于 HTTP 协议和 JSON-RPC 编码的 RPC 服务端

```go
type rpcServer struct{}

func (s *rpcServer) Hello(request string, reply *string) error {
 *reply = "hello, " + request
 return nil
}

const PROT = ":9000"

func main() {
 //注册服务：使用 rpc.RegisterName 方法注册一个名为 "rpcServer" 的服务，并将其关联到一个实现了对应方法的结构体 rpcServer 上。
 err := rpc.RegisterName("rpcServer", new(rpcServer))
 if err != nil {
  fmt.Println("register error:", err)
 }
 //处理 HTTP 请求：使用 http.HandleFunc 方法为路径 "/httprpc" 注册一个处理函数，该函数用于处理 HTTP 请求。在处理函数中，创建一个 io.ReadWriteCloser 对象，将 HTTP 请求的 Body 和 ResponseWriter 封装起来，并使用 rpc.ServeRequest 方法将请求交给 RPC 服务端处理。
 http.HandleFunc("/httprpc", func(w http.ResponseWriter, r *http.Request) {
  var conn io.ReadWriteCloser = struct {
   io.Writer
   io.ReadCloser
  }{
   ReadCloser: r.Body,
   Writer:     w,
  }
  rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
 })
 //启动 HTTP 服务：使用 http.ListenAndServe 方法启动一个 HTTP 服务，监听指定的端口（PROT），等待客户端的请求。
 http.ListenAndServe(PROT, nil)
}
```

这段 Python 代码实现了一个基于 HTTP 协议和 JSON-RPC 编码的 RPC 客户端，具体含义如下：

创建请求数据：创建一个名为 request 的字典，包含三个键值对，分别为 "id"、"params" 和 "method"。其中，"id" 表示请求的 ID，用于标识该请求；"params" 表示请求的参数，这里为一个字符串列表，只有一个元素 "pengze"；"method" 表示请求的方法，这里为 "rpcServer.Hello"，表示调用名为 "Hello" 的方法，该方法属于名为 "rpcServer" 的 RPC 服务。

发送请求数据：使用 requests.post 方法向指定的 RPC 服务端地址和端口号（"localhost:9000/httprpc"）发送一个 POST 请求，将 request 编码为 JSON 格式，并设置请求头的 Content-Type 为 "application/json"。

获取响应数据：使用 response.json() 方法获取响应数据，并将 JSON 格式的数据解码为 Python 对象。然后使用 print() 方法将其输出。

```python
request = {
    "id": 1,
    "params": ["pengze"],
    "method": "rpcServer.Hello"
}

response = requests.post("http://localhost:9000/httprpc", json=request)
print(response.json())
```
