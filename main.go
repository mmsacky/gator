package main

import (
	"log"
	"os"

	"github.com/mmsacky/gator/internal/config"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Critical error opening config: %v", err)
	}

	myState := state{
		ConfigPTR: &cfg,
	}

	myCommands := commands{
		commandMap: make(map[string]func(*state, command) error),
	}

	myCommands.register("login", handlerLogin)

	userArgsCount := len(os.Args)

	if userArgsCount < 2 {
		log.Fatalln("Error occurred: not enough arguments")
	}

	myCommand := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	err = myCommands.run(&myState, myCommand)
	if err != nil {
		log.Fatalln("Error occurred:", err)
	}
}
