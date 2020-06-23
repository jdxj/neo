package client

import (
	"net/http"
	"net/http/cookiejar"
)

func NewClient() *Client {
	jar, _ := cookiejar.New(nil)
	hClient := &http.Client{}
	hClient.Jar = jar

	c := &Client{
		hClient: hClient,
	}
	return c
}

type Client struct {
	hClient *http.Client
}

func (c *Client) GetUsageAmount() []*UsageAmount {
}

type UsageAmount struct {
	Day    string
	Amount string
}
