package goadfly

// ShortenLink - shorten new link
func (c *Client) ShortenLink(url string) (string, error) {
	endpoint := apiURL + "shorten"
	requestData := map[string]interface{}{
		"url": url,
	}
	c.sendPostRequest(endpoint, requestData)

	return "", nil // TODO
}
