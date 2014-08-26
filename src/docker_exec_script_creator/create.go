package docker_exec_script_creator

import (
	"fmt"
	parser "./../server_spec_parser"
)

func Create(spec parser.ServerSpec) []string {
	var cmds []string

    for _, container := range spec.Containers {
    	if(container.BuildRequired == true) {
			cmds = append(cmds, build_build_cmd(container))
    	}

		cmds = append(cmds, build_run_cmd(container))
	}

	return cmds
}


func build_build_cmd(c parser.Container) string {
	return fmt.Sprintf("docker build -t %s %s", c.Name, c.Image)
}

func build_run_cmd(c parser.Container) string {
	return fmt.Sprintf("docker run -d --name %s %s", c.Name, c.Image)
}