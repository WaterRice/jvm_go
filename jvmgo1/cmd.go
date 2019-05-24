package main

import (
	"fmt"
	"flag"
	"os"
)

// Format: java [-options] class [args...]
type Cmd struct {
	helpFlag      bool     // 'java --help'
	vFlag 	      bool     // 'java --version'
	classPath     string   // 'java -classpath=xxx'
	className     string   // 'java xxx.class'
	args          []string // 'java xxx.class -xxx=jjj'
}

func parseCmd() (cmd *Cmd) {
	cmd = &Cmd{}

	flag.Usage = printUsg

	// Register help flag
	// Formats: --help or --h or ?
	helpOpts :=[...]string{"help", "h", "?"}
	helpUsgMsg := "print help message."
	for _, opt := range helpOpts {
		flag.BoolVar(&cmd.helpFlag, opt, false, helpUsgMsg)
	}

	// Register version flag
	// Formats: --version or --v or --V or --Version or --VERSION
	verOpts :=[...]string{"version", "v", "V", "Version", "VERSION"}
	verUsgMsg := "print version message."
	for _, opt := range verOpts {
		flag.BoolVar(&cmd.vFlag, opt, false, verUsgMsg)
	}

	// Register classpath flag
	// Formats: -classpath xxx or -cp xxx or -ClassPath xxx
	cpOpts :=[...]string{"classpath", "cp", "ClassPath"}
	cpUsgMsg := "classpath"
	for _, opt := range cpOpts {
		flag.StringVar(&cmd.classPath, opt, "", cpUsgMsg)
	}

	// Start parse
	flag.Parse()

	// Get lefted cmd.className and cmd.args
	args := flag.Args()
	argsLength := len(args)
	if argsLength > 0 {
		cmd.className = args[0]
	}
	if argsLength > 1 {
		cmd.args = args[1:]
	}

	return
}

func printUsg() {
	fmt.Printf("Usage: %s [-option] class [args...]\n",os.Args[0])
	flag.PrintDefaults()
}