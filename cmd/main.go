package main

import (
	"log"

	"github.com/Mohs9n/chat/internal"
)

func main() {
	log.Fatal(internal.NewRouter().Run(":8080"))
}