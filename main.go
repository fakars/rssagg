package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT env not found")
	}

	fmt.Println(port)
}
