package main

import (
	"log"

	"github.com/becker63/simtricky/backend/services/agent/commands"
)

func main() {
	log.Println("Starting agent...")
	commands.Start()
}
