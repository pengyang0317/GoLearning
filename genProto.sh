function genProto {
    DOMAIN=$1
    FileAddress=$2
    # SKIP_GATEWAY=$2
    PROTO_PATH=./src/$FileAddress/proto
    GO_OUT_PATH=./src/$FileAddress/proto
    mkdir -p $GO_OUT_PATH
    
    protoc -I=$PROTO_PATH --go_out=paths=source_relative:$GO_OUT_PATH ${DOMAIN}.proto
    protoc -I=$PROTO_PATH --go-grpc_out=paths=source_relative:$GO_OUT_PATH ${DOMAIN}.proto
   
}

# genProto hello  ch4/v2
# genProto stream  ch4/v3



function genPzShopServer {
    DOMAIN=$1
    FileAddress=$2
    PROTO_PATH=./pz-shop-rpc/$FileAddress/proto
    GO_OUT_PATH=./pz-shop-rpc/$FileAddress/proto

    GO_OUT_API_PATH="./pz-shop-api/$DOMAIN-web/proto"
    protoc -I=$PROTO_PATH --go_out=paths=source_relative:$GO_OUT_PATH ${DOMAIN}.proto
    protoc -I=$PROTO_PATH --go-grpc_out=paths=source_relative:$GO_OUT_PATH ${DOMAIN}.proto
    # api处调用
    mkdir -p $GO_OUT_API_PATH
    protoc -I=$PROTO_PATH --go_out=paths=source_relative:$GO_OUT_API_PATH ${DOMAIN}.proto
    protoc -I=$PROTO_PATH --go-grpc_out=paths=source_relative:$GO_OUT_API_PATH ${DOMAIN}.proto
}
genPzShopServer user  user-srv