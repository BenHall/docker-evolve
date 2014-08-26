package docker_exec_script_creator

import (
	"testing"
	"github.com/stretchr/testify/assert"
	parser "./../server_spec_parser"
)


func TestCreate_Returns_Buld_And_Run_Command_For_Simple_Node(t *testing.T) {
	container := parser.Container {
		Name: "test-node-1",
		Image: "test/node",
		BuildRequired: true,
	}

	serverSpec := parser.ServerSpec{
		Containers: []parser.Container { container },
	}

	cmds := Create(serverSpec)
	assert.Equal(t, len(cmds), 2)
	assert.Equal(t, cmds[0], "docker build -t test-node-1 test/node")
	assert.Equal(t, cmds[1], "docker run -d --name test-node-1 test/node")
}

