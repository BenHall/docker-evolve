package docker_exec_script_creator

import (
	"testing"
	"github.com/stretchr/testify/assert"
	parser "./../server_spec_parser"
)


func TestCreate_Returns_Build_And_Run_Command_For_Simple_Node(t *testing.T) {
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


func TestCreate_Returns_Run_Command_With_Linked_Containers_Mapped_Volumes(t *testing.T) {
	container := parser.Container {
		Name: "nginx",
		Image: "dockerfile/nginx",
		Ports: []parser.Port { parser.Port {Guest: 80, Expose: 8080}},
		Links: []string { "web-app-1" },
		MappedVolumes: []parser.MappedVolume { parser.MappedVolume {Guest: "/data", Host: "/Users/root/Desktop/nginx/www"}},
	}

	serverSpec := parser.ServerSpec{
		Containers: []parser.Container { container },
	}

	cmds := Create(serverSpec)
	assert.Equal(t, len(cmds), 1)
	assert.Equal(t, cmds[0], "docker run -d --name nginx --link web-app-1:web-app-1 -p 8080:80 -v /Users/root/Desktop/nginx/www:/data dockerfile/nginx")
}
