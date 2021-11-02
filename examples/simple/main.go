package main

import (
	"log"

	"github.com/saltbo/bogo"
)

func main() {
	s, err := bogo.NewApp("config.toml")
	if err != nil {
		log.Fatalln(err)
	}

	s.Run()
}
