package main

import (
	"fmt"
	"os/exec"
    "os"
)

func Execute(cmds []string, debug bool) bool {
	file, _ := WriteLinesToTempFile(cmds);
	chmodFile(file)

	var app string
	var args string
	if(debug == true) {
		app = "cat"
		args = file
	} else {
		app = file
		args = ""
	}

	stdout, err := exec.Command(app, args).Output()
	if (err != nil) {
	   fmt.Fprintln(os.Stderr, err)
	   return false
	}

	fmt.Printf("%s\n", stdout)

    return true
}

func chmodFile(file string) {
	exec.Command("chmod", "777 " + file);
}