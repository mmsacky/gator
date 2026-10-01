package main

import (
	"fmt"
	"log"

	"github.com/mmsacky/gator/internal/config"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Critical error opening config: %v", err)
	}

	err = cfg.SetUser("Mike")
	if err != nil {
		log.Fatalf("Critical writing to config file: %v", err)
	}

	updatedCFG, err := config.Read()
	if err != nil {
		log.Fatalf("Critical error opening updated config: %v", err)
	}

	fmt.Println(updatedCFG)
}
