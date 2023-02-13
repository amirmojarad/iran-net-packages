package main

import (
	"internet_service/transport/server/http"
	"log"
)

func main() {
	server := http.Server()
	err := server.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
