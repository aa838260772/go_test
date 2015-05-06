package main

import (
    "log"
    "fmt"
    "github.com/golang/protobuf/p2p"
    "github.com/golang/protobuf/proto"
)

func main() {
    msg := &p2p.UserInfo{
        Uid: proto.Uint64(1231123),
        Type: proto.Uint32(12),
    }
    data, err := proto.Marshal(msg)
    if err != nil{
        log.Fatal("marshaling error:",err)
    }
    fmt.Printf("%v\n", data)

    new_msg := &p2p.UserInfo{}
    err = proto.Unmarshal(data, new_msg)
    if err != nil{
        log.Fatal("Unmashaling error:", err)
    }

    fmt.Printf("%v\n", new_msg)
}
