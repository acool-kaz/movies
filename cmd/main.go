package main

import (
	"log"

	"github.com/acool-kaz/movies/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
