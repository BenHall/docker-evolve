package nginx

import (
    "encoding/json"
    "strconv"
    parser "./../../server_spec_parser"
    "path/filepath"
)

type NgnixConfigSection struct {
	Format string `json:"format"`
	Target string `json:"target"`
	Domain string `json:"domain"`
	Ports []Port `json:"ports"`
}

type Port struct {
    Guest int `json:"guest"`
    Expose int `json:"expose"`
}

func Create(c parser.Container, config_json json.RawMessage) []string {
    var config NgnixConfigSection
	json.Unmarshal(config_json, &config)

	cmds := CreateCmdsForConfigFile(c, config)

	return cmds
}

func CreateCmdsForConfigFile(c parser.Container, config NgnixConfigSection) []string {
	//TODO: HORRIBLE CODE. Sorry. Only supports one port at the moment. One problem at a time.
	file := FindMappedSitesAvailable(c, config)
	port := config.Ports[0]
	guest := strconv.Itoa(port.Guest)
	expose := strconv.Itoa(port.Expose)
	upstream_name := config.Target + "-" + guest

	c1 := "echo \"upstream " + upstream_name + " { \" >> " + file
	c2 := "echo \"    server " + config.Target + ":" + guest + "; \" >> " + file
	c3 := "echo \"} \" >> " + file
	c4 := "echo \"server { \" >> " + file
	c5 := "echo \"  listen " + expose + " default_server; \" >> " + file
	c6 := "echo \"  listen [::]:" + expose + " default_server ipv6only=on; \" >> " + file
	c7 := "echo \"  server_name " + config.Domain + "; \" >> " + file
	c8 := "echo \"  location / { \" >> " + file
	c9 := "echo \"	  proxy_set_header X-Real-IP \\$remote_addr; \" >> " + file
	c10 := "echo \"    proxy_set_header X-Forwarded-For \\$proxy_add_x_forwarded_for; \" >> " + file
	c11 := "echo \"    proxy_set_header Host \\$http_host; \" >> " + file
	c12 := "echo \"    proxy_set_header X-NginX-Proxy true; \" >> " + file
	c13 := "echo \"    proxy_pass http://" + upstream_name + "; \" >> " + file
	c14 := "echo \"    proxy_redirect off; \" >> " + file
	c15 := "echo \"  } \" >> " + file
	c16 := "echo \"} \" >> " + file

	return []string{ c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12, c13, c14, c15, c16}
}

func FindMappedSitesAvailable(c parser.Container, config NgnixConfigSection) string {
	file := config.Target + "-nginx.conf"

	var dir string
	for _, m := range c.MappedVolumes {
		if(m.Guest == "/etc/nginx/sites-enabled") {
			dir = m.Host
		}
	}

	return filepath.Join(dir, file)
}


