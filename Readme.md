# Docker Evolve

Evolve a server from nothing to multiple Docker containers running in your desired configuration in a single command

## Aim of Docker Evolve

* Ability to place host configuration for Docker under source control
* Maintain container links, volumes and ports for a host
* Deploy changes to a host configuration in a repeatable fashion
* Isolated command line tool capable of evolving either an individual or cluster of servers.

## Usage Example

Based on the server_node_nginx.json example. Command will build and run a docker container for a node.js application with an nginx container making the node.js application accessible.

```
./docker-evolve examples/server_node_nginx.json
```

## Debug

Set the debug flag to see the output commands to STDOUT instead of executing

```
./docker-evolve -debug=true examples/server_node_nginx.json
```

