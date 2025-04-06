package elgatoring

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	DefaultPort uint16 = 9123
)

type Client struct {
	c    *http.Client
	port uint16
	host string
}

func New(host string, opts ...ClientOption) (*Client, error) {
	h, p, err := parseHost(host)
	if err != nil {
		return nil, err
	}

	if p == 0 {
		p = DefaultPort
	}

	c := &Client{
		c:    http.DefaultClient,
		host: h,
		port: p,
	}

	return c, nil
}

func (c *Client) request(
	ctx context.Context,
	method string,
	path string,
	body any,
	result any,
) error {
	url := fmt.Sprintf("http://%s:%d%s", c.host, c.port, path)
	var r io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return err
		}
		r = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, r)
	if err != nil {
		return err
	}

	if r != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	res, err := c.c.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("status code: %d", res.StatusCode)
	}

	if result == nil {
		return nil
	}

	return json.NewDecoder(res.Body).Decode(result)
}

func (c *Client) GetAccessoryInfo(ctx context.Context) (*AccessoryInfo, error) {
	var info AccessoryInfo

	if err := c.request(ctx, http.MethodGet, "/elgato/accessory-info", nil, &info); err != nil {
		return nil, err
	}

	return &info, nil
}

func (c *Client) GetLights(ctx context.Context) ([]*Light, error) {
	var data struct {
		Count  int      `json:"numberOfLights"`
		Lights []*Light `json:"lights"`
	}

	if err := c.request(ctx, http.MethodGet, "/elgato/lights", nil, &data); err != nil {
		return nil, err
	}

	return data.Lights, nil
}

func (c *Client) SetLights(ctx context.Context, lights []*Light) ([]*Light, error) {
	data := struct {
		Count  int      `json:"numberOfLights"`
		Lights []*Light `json:"lights"`
	}{
		Count:  len(lights),
		Lights: lights,
	}

	if err := c.request(ctx, http.MethodPut, "/elgato/lights", &data, &data); err != nil {
		return nil, err
	}

	return data.Lights, nil
}

func (c *Client) Identify(ctx context.Context) error {
	return c.request(ctx, http.MethodPost, "/elgato/identify", nil, nil)
}
