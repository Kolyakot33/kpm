package main

import (
	"log"
	"net/rpc"
)

func sendSignal(name string, args ...string) error {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:7124")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply int
	err = client.Call("KPMD."+name, args, &reply)
	if err != nil {
		println(err.Error())
		return err
	}
	return nil

}

/*
func getLogFile(id string) string {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:7124")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply string
	err = client.Call("KPMD.Log", id, &reply)
	if err != nil {
		log.Fatal(err.Error())
	}
	return reply

}*/

// Pulls stdout and stderr from a process and returns them as a string.
func askStdOut(id string) string {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:7124")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply string
	err = client.Call("KPMD.StdOut", id, &reply)
	if err != nil {
		log.Fatal(err.Error())
	}
	return reply
}

func reqList() error {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:7124")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply []ProcessInfo
	err = client.Call("KPMD.List", "", &reply)
	if err != nil {
		println(err.Error())
		return err
	}
	println("------------------------------")
	println("| ID |\tNAME\t|\tARGS\t|\tSTATUS\t|\tPID\t|")
	for _, process := range reply {
		println("------------------------------")
		var args string
		for _, arg := range process.Args {
			args += arg + " "
		}
		println("|", process.Id, "\t|", process.File, "\t|", args, "\t|", process.State, "\t|", process.Pid, "\t|")
	}
	println("------------------------------")
	return nil
}
