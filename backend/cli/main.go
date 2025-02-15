package main

import (
	"log"

	"github.com/becker63/simtricky/services/agent/commands"
)

func main() {
	log.Println("Starting agent...")
	commands.Start()
}
