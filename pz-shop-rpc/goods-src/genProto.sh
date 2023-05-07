function genServer {
    DOMAIN=$1
    PROTO_PATH=./proto
    GO_OUT_PATH=$PROTO_PATH
    protoc -I=$PROTO_PATH --go_out=paths=source_relative:$GO_OUT_PATH ${DOMAIN}.proto
    protoc -I=$PROTO_PATH --go-grpc_out=paths=source_relative:$GO_OUT_PATH ${DOMAIN}.proto
}
genServer goods