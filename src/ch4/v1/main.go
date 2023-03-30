package main

import (
	"encoding/json"
	"fmt"
	hellopb "lgo/src/ch4/v1/proto"

	"google.golang.org/protobuf/proto"
)

type Hello struct {
	Name string `json:"name"`
}

func main() {

	req := hellopb.HelloRequest{
		Name: "pengze",
	}

	by, _ := proto.Marshal(&req)

	jsonStrut := Hello{Name: "pengze"}
	jsonReq, _ := json.Marshal(jsonStrut)

	fmt.Printf("%v,%v\n", by, string(by))
	fmt.Printf("%v,%v", jsonReq, string(jsonReq))
}
