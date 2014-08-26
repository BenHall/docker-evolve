package main

import (
    "fmt"
    "os"
    "flag"
    Parser "./server_spec_parser"
    Docker "./docker_exec_script_creator"
)

var serverConfigFile = flag.String("config", "", "Server Configuration File")
var serverTargetIpAddress = flag.String("ip", "", "IP of Target Server")
var debug = flag.Bool("debug", false, "Debug Flag - Docker commands will be outputted instead of executed")

func main() {
    flag.Parse()

    if(*serverConfigFile == "") {
        if(len(flag.Args()) == 0) {
            fmt.Fprintf(os.Stderr, "Error: No server configuration provided \n")
            os.Exit(1)
        } else {
            serverConfigFile = &flag.Args()[0]
        }
    }

    serverSpec := Parser.ServerSpec{}
    err := serverSpec.Parse(*serverConfigFile)
    if err != nil { panic(err) }


    cmds := Docker.Create(serverSpec)
    Execute(cmds, *debug)
    os.Exit(0)
}