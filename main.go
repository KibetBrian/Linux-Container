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

func main() {
	log.Println("Arguments: ", os.Args)

	if len(os.Args) < 2 {
		log.Println("Check commands, try: linux-container")
		os.Exit(1)
	}

	switch os.Args[1] {
		case "linux-container":
		InitProcess()

		case "run":

			if os.Args[0] != "/proc/self/exe"{
				log.Println("Invalid command, try: linux-container")
				os.Exit(1)
			}
			StartMainProcess()

		default:
		log.Println("Invalid commands, try: linux-container")
		os.Exit(1)
	}
}

func InitProcess() {
	log.Println("ProcessId: ", os.Getpid())
	arguments := []string{"/proc/self/exe", "run"}

	cmd := exec.Command(arguments[0], arguments[1:]...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	cmd.SysProcAttr = NameSpacesFlags(cmd)

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func StartMainProcess() {
	setHostName()
	arguments := []string{"/bin/bash"}
	
	cmd := exec.Command(arguments[0])
	MountFileSystem()

	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	cmd.SysProcAttr = NameSpacesFlags(cmd)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error: %v \n", err)
		os.Exit(1)
	}
	err = syscall.Unmount("/proc",0)
	if err != nil {
		panic(err)
	}
	log.Printf("Exited the container with process id: %v \n", os.Getpid())
}

func NameSpacesFlags(cmd *exec.Cmd) *syscall.SysProcAttr {
	attributes := &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWPID |
		 			syscall.CLONE_NEWUTS|
					syscall.CLONE_NEWIPC|
		 			syscall.CLONE_NEWNS,
	}
	return attributes
}

func MountFileSystem (){
	err := syscall.Chroot("/home/root")
	if err != nil {
		panic(err)
	}

	err = syscall.Chdir("/")	
	if err != nil {
		panic(err)
	}
	
	err = syscall.Mount("proc", "proc", "proc", 0, "")
	if err != nil {
		panic(err)
	}
}

func setHostName() {
	err := syscall.Sethostname([]byte(hostname))
	if err != nil {
		panic(err)
	}
}
