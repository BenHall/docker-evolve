package server_spec_parser

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestParse_Parses_Node_Container(t *testing.T) {
	serverSpec := &ServerSpec{}
    serverSpec.Parse("./../../examples/server_node_nginx.json")
    c := serverSpec.Containers[0]

	assert.Equal(t, c.Name, "node-app-1")
	assert.Equal(t, c.Image, "github.com/BenHall/docker-node-example")
	assert.Equal(t, c.BuildRequired, true)
}
