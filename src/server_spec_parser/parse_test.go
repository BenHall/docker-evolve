package server_spec_parser

import (
	"testing"
	"github.com/stretchr/testify/assert"
    "encoding/json"
)


func TestParse_Parses_Node_Container(t *testing.T) {
	serverSpec := &ServerSpec{}
    serverSpec.Parse("./../../examples/server_node_nginx.json")
    c := serverSpec.Containers[0]

	assert.Equal(t, c.Name, "node-app-1")
	assert.Equal(t, c.Image, "github.com/BenHall/docker-node-example")
	assert.Equal(t, c.BuildRequired, true)
}

func TestParse_Parses_Config_Section(t *testing.T) {
	serverSpec := &ServerSpec{}
    serverSpec.Parse("./../../examples/server_node_nginx.json")
    c := serverSpec.Containers[1]

	assert.NotEqual(t, c.Config, "")
	assert.Equal(t, len(c.Config), 1)

	var objmap map[string]*json.RawMessage
	json.Unmarshal(c.Config[0], &objmap)

	var str string
	json.Unmarshal(*objmap["format"], &str)

	assert.Equal(t, str, "nginx")
}

func TestGetConfigFormat_ReturnsString_For_Config(t *testing.T) {
	serverSpec := &ServerSpec{}
    serverSpec.Parse("./../../examples/server_node_nginx.json")
    c := serverSpec.Containers[1]
    str := c.GetConfigFormat(c.Config[0])

	assert.Equal(t, str, "nginx")
}





