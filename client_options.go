package elgatoring

import (
	"net/http"
	"regexp"
	"strconv"
)

var portPattern = regexp.MustCompile(`:(\d+)$`)

type ClientOption func(*Client) error

func WithHTTPClient(c *http.Client) ClientOption {
	return func(client *Client) error {
		client.c = c
		return nil
	}
}

func WithHost(host string) ClientOption {
	return func(client *Client) error {
		h, p, err := parseHost(host)
		if err != nil {
			return err
		}

		client.host = h
		if p != 0 {
			client.port = p
		}

		return nil
	}
}

func parseHost(host string) (string, uint16, error) {
	m := portPattern.FindStringSubmatch(host)
	if m == nil {
		return host, 0, nil
	}

	p, err := strconv.ParseUint(m[1], 10, 16)
	if err != nil {
		return "", 0, err
	}

	return host[:len(host)-len(m[0])], uint16(p), nil
}
