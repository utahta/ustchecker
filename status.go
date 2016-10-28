package uststat

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/mattn/go-scan"
)

type Client struct {
	httpClient *http.Client
}

type ClientOption func(*Client) error

// New client
func New(options ...ClientOption) (*Client, error) {
	c := &Client{httpClient: http.DefaultClient}

	for _, option := range options {
		if err := option(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

// WithHTTPTransport function
func WithHTTPTransport(t http.RoundTripper) ClientOption {
	return func(c *Client) error {
		c.httpClient.Transport = t
		return nil
	}
}

// Get ustream live status by channel name
func (c *Client) IsLive(name string) (bool, error) {
	resp, err := c.httpClient.Get(fmt.Sprintf("http://www.ustream.tv/channel/%s", name))
	if err != nil {
		return false, err
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return false, err
	}

	id, ok := doc.Find("meta[name='ustream:channel_id']").Attr("content")
	if !ok {
		return false, errors.New("Failed to get ustream:channel_id")
	}
	return c.IsLiveByChannelID(id)
}

// Get ustream live status by channel id
func (c *Client) IsLiveByChannelID(id string) (bool, error) {
	resp, err := c.httpClient.Get(fmt.Sprintf("https://api.ustream.tv/channels/%s.json", id))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var status string
	if err := scan.ScanJSON(resp.Body, "/channel/status", &status); err != nil {
		return false, err
	}
	return status == "live", nil
}
