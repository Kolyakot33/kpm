package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

const version = "0.0.1"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}
	switch os.Args[1] {
	case "daemon":
		runDaemon()
	case "start":
		wd, _ := os.Getwd()
		sendSignal("Run", append(os.Args[1:], wd)...)
	case "stop":
		sendSignal("Stop", os.Args[1:]...)
	case "list":
		reqList()
	case "restart":
		sendSignal("Restart", os.Args[1:]...)
	case "kill":
		sendSignal("Kill", os.Args[1:]...)
	case "help", "--help", "-h":
		printUsage()
	}
}

// Prints help and usage information to stdout.
func printUsage() {
	fmt.Printf("KPMD v%s - A simple process manager\n", version)
	println("Author: Kolyakot33 https://l0sty.ru")
	println("Usage: kpmd [command] [args]")
	println("Commands:")
	println("\tstart [file] [args] - Starts a new process with the given file and args.")
	println("\tstop [id] - Stops the process with the given id.")
	println("\trestart [id] - Restarts the process with the given id.")
	println("\tkill [id] - Kills the process with the given id.")
	println("\tlist - Lists all running processes.")
	println("\tdaemon - Starts a new daemon process.")
	println("\thelp - Prints this help message.")
}

func runDaemon() {
	cmd := exec.Command("kpmd")

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	if cmd.Process == nil {
		println("Failed to start kpmd")
	}
}

type ProcessInfo struct {
	Id, Pid     int
	File, State string
	Args        []string
}
