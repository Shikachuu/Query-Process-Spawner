package main

import (
	"github.com/Shikachuu/php-process-redis-list/cmd"
	"log"
)

func main() {
	c := cmd.RootCommand()
	err := c.Execute()
	if err != nil {
		log.Fatalf("Failed to start runner. Error: %s", err)
	}
}
