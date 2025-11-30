package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	interactive := flag.Bool("interactive", false, "Run in interactive mode")
	mode := flag.String("mode", "notify", "Execution mode: reserve, notify, or freeze")
	configFile := flag.String("config", ".env.local", "Path to configuration file")

	flag.Parse()

	if err := godotenv.Load(*configFile); err != nil {
		log.Printf("Warning: Could not load %s: %v", *configFile, err)
	}

	fmt.Println("\nRuntime Parameters:")
	fmt.Println("Interactive mode:", *interactive)
	fmt.Println("Execution mode:", *mode)
	fmt.Println("Config file:", *configFile)

	fmt.Println("\nConfiguration:")
	fmt.Println("BASE_URL:", os.Getenv("BASE_URL"))
	fmt.Println("USERNAME:", os.Getenv("USERNAME"))

	if *mode != "" {
		validModes := map[string]bool{"reserve": true, "notify": true, "freeze": true}
		if !validModes[*mode] {
			log.Fatalf("Invalid mode: %s. Valid modes are: reserve, notify, freeze", *mode)
		}
	}

	result := greet("Padel Player")
	fmt.Println(result)
}

func greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
