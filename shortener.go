package goadfly

import (
	"log"
	"time"
)

// ShortenLink - shorten new link
func (c *Client) ShortenLink(url string) (string, error) {
	endpoint := apiURL + "shorten"
	requestData := map[string]interface{}{
		"_user_id":   c.UserID,
		"_api_key":   c.APIKeyPublic,
		"_timestamp": time.Now().Unix(),
		"url":        url,
	}
	response, err := c.sendPostRequest(endpoint, requestData)
	if err != nil {
		return "", err
	}

	log.Println(string(response)) // TEMP

	return "", nil // TODO
}
