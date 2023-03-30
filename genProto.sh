function genProto {
    DOMAIN=$1
    SKIP_GATEWAY=$2
    PROTO_PATH=./src/ch4/v1
    GO_OUT_PATH=./src/ch4/v1
    # mkdir -p $GO_OUT_PATH
    
    protoc -I=$PROTO_PATH --go_out=paths=source_relative:$GO_OUT_PATH ${DOMAIN}.proto
   
}

genProto hello