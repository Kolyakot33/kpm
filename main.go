package main

import (
	"fmt"
	"golang.org/x/term"
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
		reqList()

	case "start":
		wd, _ := os.Getwd()
		sendSignal("Run", append(os.Args[2:], wd)...)
		reqList()

	case "stop":
		sendSignal("Stop", os.Args[2:]...)
		reqList()

	case "list":
		reqList()
	case "restart":
		sendSignal("Restart", os.Args[2:]...)
		reqList()

	case "kill":
		sendSignal("Kill", os.Args[2:]...)
		reqList()
	case "attach":
		attach(os.Args[2])

	case "help", "--help", "-h":
		printUsage()
	}
}

// Prints help and usage information to stdout.
func printUsage() {
	fmt.Printf("KPM v%s - A simple process manager\n", version)
	fmt.Println("Author: Kolyakot33 https://l0sty.ru")
	fmt.Println("Usage: kpm [command] [args]")
	fmt.Println("Commands:")
	fmt.Println("\tstart [file] [args] - Starts a new process with the given file and args.")
	fmt.Println("\tstop [id] - Stops the process with the given id.")
	fmt.Println("\trestart [id] - Restarts the process with the given id.")
	fmt.Println("\tkill [id] - Kills the process with the given id.")
	fmt.Println("\tattach [id] - Attaches to the process with the given id.")
	fmt.Println("\tlist - Lists all running processes.")
	fmt.Println("\tdaemon - Starts a new daemon process.")
	fmt.Println("\thelp - Prints this help message.")
}

type sh struct{}

func (sh *sh) Read(b []byte) (int, error) {
	return os.Stdin.Read(b)
}

func (sh *sh) Write(b []byte) (int, error) {
	return os.Stdout.Write(b)
}

func attach(id string) {
	_, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	//defer term.Restore(int(os.Stdin.Fd()), oldState)
	terminal := term.NewTerminal(&sh{}, ">")
	go func() {
		for {
			line, err := terminal.ReadLine()
			if err != nil {
				log.Fatal(err)
			}
			sendSignal("Stdin", []string{id, line}...)
		}
	}()

	func() {
		for {
			terminal.Write([]byte(askStdOut(id) + "\n"))
		}
	}()

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
