package main

import (
	"fmt"
)

func main() {
	cmd := parseCmd()
	if cmd.vFlag {
		fmt.Println("MyJVM built via go, version 0.0.1")
	} else if cmd.helpFlag || cmd.className == "" {
		printUsg()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	fmt.Printf("classpath:%s\tclassName:%s\targs:%s\t\n", 
		cmd.classPath, cmd.className, cmd.args)
}

