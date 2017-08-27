package main 

import (
	"fmt"
	"os"
	_ "os/exec"
	_ "syscall"

)

func main(){
	switch os.Args[1] {
	case "host" : 
		host()
	case "container":
		// container()
	default:
		fmt.Println("Usage: barebonescontainer <mode>")
	}
}

func host(){
	pid := os.Getpid()
	fmt.Println("pid", pid)
}