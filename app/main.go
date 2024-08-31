package main

import (
	"fmt"
	"log"

	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client"
)

func main() {
	// Create a new client
	client := artifactsmmo.NewClient(nil)

	// Example: Get player information
	player, _, err := client.Players.Get("your-player-id")
	if err != nil {
		log.Fatalf("Error getting player information: %v", err)
	}

	fmt.Printf("Player Name: %s\n", player.Name)
	fmt.Printf("Player Level: %d\n", player.Level)
}
