package uststat

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PuerkitoBio/goquery"
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

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var data interface{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		return false, err
	}
	root := data.(map[string]interface{})
	channel := root["channel"].(map[string]interface{})
	status := channel["status"].(string)

	return status == "live", nil
}
