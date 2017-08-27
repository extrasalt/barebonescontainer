package main 

import (
	"fmt"
	"os"
	"os/exec"
	_ "syscall"

)

func main(){
	switch os.Args[1] {
	case "host" : 
		host()
	case "container":
		container()
	default:
		fmt.Println("Usage: barebonescontainer <mode>")
	}
}

func host(){
	pid := os.Getpid()
	fmt.Println("pid", pid)
}

func container(){
	cmd := exec.Command("/proc/self/exe", "host")
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		panic(err)
	}
}