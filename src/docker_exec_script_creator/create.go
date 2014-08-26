package docker_exec_script_creator

import (
	"fmt"
	parser "./../server_spec_parser"
	nginx "./../config_creators/nginx"
)

func Create(spec parser.ServerSpec) []string {
	var cmds []string

    for _, container := range spec.Containers {
    	//TODO: This needs to be better. Not flexible enough to support build + config
    	if(container.BuildRequired == true) {
			cmds = append(cmds, build_build_cmd(container))
			container.Image = container.Name
			cmds = append(cmds, build_run_cmd(container))
    	} else {

			if(len(container.Config) > 0) {
				config_cmds := nginx.Create(container, container.Config[0])
				cmds = append(cmds, config_cmds...)
			}
			
			cmds = append(cmds, build_run_cmd(container))
    	}

	}

	return cmds
}

func build_build_cmd(c parser.Container) string {
	return fmt.Sprintf("docker build -t %s %s", c.Name, c.Image)
}

func build_run_cmd(c parser.Container) string {
	name := fmt.Sprintf("--name %s ", c.Name)
	var links string
	var ports string
	var mapped_volumes string
	var environment string

	if(len(c.Links) > 0) {
		links = create_links_section(c.Links)
	}
	if(len(c.Ports) > 0) {
		ports = create_ports_section(c.Ports)
	}	
	if(len(c.Ports) > 0) {
		mapped_volumes = create_mapped_volumes_section(c.MappedVolumes)
	}
	if(len(c.EnvironmentVariables) > 0) {
		environment = create_environment_section(c.EnvironmentVariables)
	}

	cmd := fmt.Sprintf("docker run -d %s%s%s%s%s%s", name, links, ports, mapped_volumes, environment, c.Image)

	return cmd
}

func create_links_section(links []string) string {
    var cmd string

    for _, l := range links {
    	cmd = fmt.Sprintf("%s--link %s:%s ", cmd, l, l)
	}

	return cmd;
}

func create_ports_section(ports []parser.Port) string {
    var cmd string

    for _, p := range ports {
    	cmd = fmt.Sprintf("%s-p %d:%d ", cmd, p.Expose, p.Guest)
	}

	return cmd;
}

func create_mapped_volumes_section(mapped_volumes []parser.MappedVolume) string {
    var cmd string

    for _, m := range mapped_volumes {
    	cmd = fmt.Sprintf("%s-v %s:%s ", cmd, m.Host, m.Guest)
	}

	return cmd;
}

func create_environment_section(mapped_volumes []parser.EnvironmentVariable) string {
    var cmd string

    for _, m := range mapped_volumes {
    	cmd = fmt.Sprintf("%s-e %s=%s ", cmd, m.Key, m.Value)
	}

	return cmd;
}




