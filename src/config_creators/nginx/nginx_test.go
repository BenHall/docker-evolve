package nginx

import (
	"testing"
	"github.com/stretchr/testify/assert"
    parser "./../../server_spec_parser"
)

func TestCreateCmdsForConfigFile_Returns_Cmds_To_Create_Config_Script(t *testing.T) {
	container := parser.Container {
		Name: "nginx",
		Image: "dockerfile/nginx",
		Ports: []parser.Port { parser.Port {Guest: 80, Expose: 8080}},
		Links: []string { "web-app-1" },
		MappedVolumes: []parser.MappedVolume { parser.MappedVolume {Guest: "/data", Host: "/Users/root/Desktop/nginx/www"}},
	}

	config := NgnixConfigSection {
		Format: "nginx",
		Target: "node-app-1",
		Domain: "localhost",
		Ports: []Port { Port { Guest: 3000, Expose: 80} },
	}

	cmds := CreateCmdsForConfigFile(container, config)
	assert.Equal(t, len(cmds), 16)
}

func TestFindMappedSitesAvailable_Returns_Volume_With_Filename(t *testing.T) {
	container := parser.Container {
		Name: "nginx",
		Image: "dockerfile/nginx",
		Ports: []parser.Port { parser.Port {Guest: 80, Expose: 8080}},
		Links: []string { "web-app-1" },
		MappedVolumes: []parser.MappedVolume { parser.MappedVolume {Guest: "/etc/nginx/sites-enabled", Host: "/Users/root/Desktop/nginx/sites-enabled"}},
	}

	config := NgnixConfigSection {
		Format: "nginx",
		Target: "node-app-1",
		Domain: "localhost",
		Ports: []Port { Port { Guest: 3000, Expose: 80} },
	}

	target := "/Users/root/Desktop/nginx/sites-enabled/node-app-1-nginx.conf"

	result := FindMappedSitesAvailable(container, config)
	assert.Equal(t, target, result)
}