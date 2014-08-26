package main

import (
    "fmt"
    Parser "./server_spec_parser"
)

func main() {
    ServerSpec := &Parser.ServerSpec{}
    err := ServerSpec.Parse("./../examples/server_node_nginx.json")
    if err != nil { panic(err) }


    fmt.Println(ServerSpec)
}