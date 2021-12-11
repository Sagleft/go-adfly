package main

import (
	"goadfly"
	"log"
)

func main() {
	client := goadfly.NewClient(0, "") // TODO
	_, err := client.ShortenLink("https://example.com")
	if err != nil {
		log.Fatalln(err)
	}
}
