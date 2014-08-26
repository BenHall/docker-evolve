package main

import (
	"os/exec"
    "os"
)

func Execute(cmds []string, debug bool) {
	file, _ := WriteLinesToTempFile(cmds);
	var app string
	var args string
	if(debug == true) {
		app = "cat"
		args = file
	} else {
		app = "/bin/bash"
		args = file
	}

	cmd := exec.Command(app, args)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()
}