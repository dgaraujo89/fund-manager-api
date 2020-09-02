package main

import (
	"fmt"
	"os"
)

const version = "0.0.1"

func main() {
	fmt.Println("Fund Manager API")
	fmt.Printf("Version: %s\n", version)

	os.Exit(0)
}
