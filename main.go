package main

import (
	"github.com/Shikachuu/Query-Process-Spawner/cmd"
	"log"
)

func main() {
	c := cmd.RootCommand()
	err := c.Execute()
	if err != nil {
		log.Fatalf("Failed to start runner. Error: %s", err)
	}
}
