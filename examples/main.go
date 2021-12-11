package main

import (
	"log"
	"os"
	"strconv"

	goadfly "github.com/Sagleft/go-adfly"
)

func main() {
	userIDRaw := os.Getenv("USER_ID")
	if userIDRaw == "" {
		log.Fatalln("user ID is not set by env")
	}

	userID, err := strconv.ParseInt(userIDRaw, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}

	apiKeyPublic := os.Getenv("API_PUBLIC")

	client := goadfly.NewClient(userID, apiKeyPublic)
	shortLink, err := client.ShortenLink("https://example.com")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(shortLink)
}
