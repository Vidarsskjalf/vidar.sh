package main

import (
	"log"
	"vidar.sh/server"
)

func main() {
	// Start server
	srv := server.NewServer(":8080")
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
