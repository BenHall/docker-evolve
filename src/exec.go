package main

import (
	"fmt"
	"os/exec"
    "os"
)

func Execute(cmds []string, debug bool) bool {
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

	stdout, err := exec.Command(app, args).Output()

	fmt.Printf("%s\n", stdout)
	if (err != nil) {
	   fmt.Printf("%s\n", err)
	   fmt.Fprintln(os.Stderr, err)
	   return false
	}


    return true
}