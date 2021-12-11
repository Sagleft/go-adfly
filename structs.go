package goadfly

type adflyLinkShortenResponse struct {
	Errors   []adflyError     `json:"errors"`
	Warnings []adflyError     `json:"warnings"`
	Data     []adflyURLResult `json:"data"`
}

type adflyError struct {
	Code     int    `json:"code"`
	Messsage string `json:"msg"`
	OnURL    string `json:"on"`
}

type adflyURLResult struct {
	ID       string `json:"id"`
	URL      string `json:"url"`
	ShortURL string `json:"short_url"`
}
