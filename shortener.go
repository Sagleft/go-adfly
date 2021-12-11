package goadfly

import (
	"encoding/json"
	"errors"
	"net/url"
)

// ShortenLink - shorten new link
func (c *Client) ShortenLink(urlToShorten string) (string, error) {
	endpoint := apiURL + "shorten"
	params := url.Values{}
	params.Add("url", urlToShorten)

	response, err := c.sendPostRequest(endpoint, params)
	if err != nil {
		return "", err
	}

	result := adflyLinkShortenResponse{}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return "", err
	}

	// find first error
	if len(result.Errors) > 0 {
		for _, firstError := range result.Errors {
			errMsg := firstError.Messsage
			if firstError.Messsage == "" {
				errMsg = "unknown error"
			}
			return "", errors.New(errMsg)
		}
	}

	for _, firstURLResult := range result.Data {
		return firstURLResult.ShortURL, nil
	}

	return "", errors.New("failed to find short url in response")
}
