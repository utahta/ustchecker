package uststat

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type mockTransport struct {
	live bool
}

func newMockTransport(live bool) http.RoundTripper {
	return &mockTransport{live: live}
}

func (t *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	response := &http.Response{
		Header:     make(http.Header),
		Request:    req,
		StatusCode: http.StatusOK,
	}
	response.Header.Set("Content-Type", "application/json")

	var status string
	if t.live {
		status = "live"
	} else {
		status = "offline"
	}
	responseBody := fmt.Sprintf(`{"channel": {"status": "%s"}}`, status)
	response.Body = ioutil.NopCloser(strings.NewReader(responseBody))
	return response, nil
}

func TestClient_IsLiveByChannelID(t *testing.T) {
	var (
		c   *Client
		b   bool
		err error
	)

	c, _ = New(WithHTTPTransport(newMockTransport(true)))
	b, err = c.IsLiveByChannelID("1234567")
	if !b || err != nil {
		t.Errorf("Expected true, got %v, %v", b, err)
	}

	c, _ = New(WithHTTPTransport(newMockTransport(false)))
	b, err = c.IsLiveByChannelID("1234567")
	if b || err != nil {
		t.Errorf("Expected false, got %v, %v", b, err)
	}
}
