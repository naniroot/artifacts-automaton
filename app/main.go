package main

import (
	"fmt"
	"os"
)
var token string


func main() {
	err := loadEnvVars()
	if err != nil {
		fmt.Println("Error loading environment variables:", err)
		return
	}

	fmt.Println(token)
}

func loadEnvVars() error {
	token = os.Getenv("ARTIFACTS_TOKEN")
	if token == "" {
		return fmt.Errorf("ARTIFACTS_TOKEN environment variable not set")
	}
	return nil
}