function genProto {
    PROTO_PATH=./user-web/protoapi
    

    protoc -I=$PROTO_PATH --go_out=paths=source_relative:$PROTO_PATH $1.proto
    protoc -I=$PROTO_PATH --go-grpc_out=paths=source_relative:$PROTO_PATH $1.proto
}

genProto userapi  user-web