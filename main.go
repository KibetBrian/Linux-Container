package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

const (
	hostname = "container"
)


func main (){
	log.Println("Arguments: ",os.Args)
	if len(os.Args)<2{
		log.Println("Check commands, try: linux-container")
		os.Exit(1)
	}
	if os.Args[1]=="linux-container" {
		InitProcess()
	}
	if os.Args[1] == "run"{
		StartMainProcess()
	}else{
		log.Println("Check commands and try again 2")
		os.Exit(1)
	}
}

func InitProcess (){
	log.Println("ProcessId: ", os.Getpid())
	arguments := []string{"/proc/self/exe", "run"}
	
	cmd := exec.Command(arguments[0], arguments[1:]...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr= os.Stdin, os.Stdout, os.Stderr
	cmd.SysProcAttr = Flags(cmd)
	
 	err := cmd.Run()
	 if err != nil {
		 panic(err)
	 }
	 log.Println("End")
}

func StartMainProcess (){
	setHostName()
	arguments := []string{"/bin/bash"}
	cmd := exec.Command(arguments[0])
	cmd.Stdin, cmd.Stdout, cmd.Stderr= os.Stdin, os.Stdout, os.Stderr
	cmd.SysProcAttr = Flags(cmd)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error: %v \n",err)
		os.Exit(1)
	}
	
}

func Flags(cmd *exec.Cmd) *syscall.SysProcAttr{
	attributes := &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWPID|syscall.CLONE_NEWUTS,
	}
	return attributes
}

func setHostName (){
	err := syscall.Sethostname([]byte(hostname))
	if err != nil {
		panic(err)
	}
}