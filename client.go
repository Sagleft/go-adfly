package goadfly

const (
	apiURL = "https://adf.ly/v1/"
)

// Client - ad.fly API client
type Client struct{}

// NewClient - create new ad.fly API client
func NewClient() *Client {
	return &Client{}
}
