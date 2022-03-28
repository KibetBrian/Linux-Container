package main

import (
	"log"
	"os"
	"os/exec"
)


func main (){

	if len(os.Args)<2{
		log.Println("Check commands, try: linux-container")
		os.Exit(1)
	}
	if os.Args[1]=="linux-container"{
		StartProcess()
	}else{
		log.Fatalln("Check commands...Too many arguments")
	}
}

func StartProcess (){

	arguments := []string{"/bin/bash"}

	cmd := exec.Command(arguments[0], arguments[1:]...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr= os.Stdin, os.Stdout, os.Stderr

 	err := cmd.Run()
	 if err != nil {
		 panic(err)
	 }
	
}