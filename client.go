package goadfly

const (
	apiURL = "https://api.adf.ly/v1/"
)

// Client - ad.fly API client
type Client struct {
	UserID       int64
	APIKeyPublic string
}

// NewClient - create new ad.fly API client
func NewClient(userID int64, apiKeyPublic string) *Client {
	return &Client{
		UserID:       userID,
		APIKeyPublic: apiKeyPublic,
	}
}
