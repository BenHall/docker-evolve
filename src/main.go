package main

import (
    "fmt"
    "os"
    "flag"
    Parser "./server_spec_parser"
)

var serverConfigFile = flag.String("config", "", "Server Configuration File")
var serverTargetIpAddress = flag.String("ip", "", "IP of Target Server")

func main() {
    flag.Parse()

    if(*serverConfigFile == "") {
        fmt.Fprintf(os.Stderr, "Error: No server configuration provided \n")
        os.Exit(1)
    }


    ServerSpec := &Parser.ServerSpec{}
    err := ServerSpec.Parse(*serverConfigFile)
    if err != nil { panic(err) }


    fmt.Println(ServerSpec)
}