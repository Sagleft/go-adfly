package goadfly

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func (c *Client) sendPostRequest(requestURL string, params url.Values) ([]byte, error) {
	params.Add("_user_id", strconv.FormatInt(c.UserID, 10))
	params.Add("_api_key", c.APIKeyPublic)

	// declare HTTP Method and Url
	resp, err := http.PostForm(requestURL, params)
	if err != nil {
		return nil, err
	}

	// read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("failed to read request response: " + err.Error())
	}

	defer resp.Body.Close()
	return body, nil
}
