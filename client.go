package go_routeros_client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ogi4i/go-routeros-client/query"
)

type Client struct {
	baseURL, username, password string
	httpClient                  http.Client
}

const defaultHTTPTimeout = 10 * time.Second

func NewClient(baseURL, username, password string, opts ...Option) *Client {
	c := Client{
		baseURL:  baseURL + "/rest",
		username: username,
		password: password,
		httpClient: http.Client{
			Transport: http.DefaultTransport,
			Timeout:   defaultHTTPTimeout,
		},
	}

	for _, o := range opts {
		o(&c)
	}

	return &c
}

func (c *Client) do(ctx context.Context, method, uri string, body, response any) error {
	var buf bytes.Buffer
	if body != http.NoBody {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			return fmt.Errorf("encode request body: %w", err)
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, uri, &buf)
	if err != nil {
		return fmt.Errorf("build request: %w", err)
	}

	req.SetBasicAuth(c.username, c.password)

	// body is always JSON encoded
	if body != http.NoBody {
		req.Header.Add("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	buf.Reset()
	if _, err = buf.ReadFrom(resp.Body); err != nil {
		return fmt.Errorf("read response body: %w", err)
	}

	if resp.StatusCode > 399 {
		var errStruct Error
		// try to decode response as error struct
		if err = json.Unmarshal(buf.Bytes(), &errStruct); err != nil {
			return fmt.Errorf("decode response body to error struct: %w", err)
		}

		return errStruct
	}

	// skip decoding body if we got a 204 response status code
	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	if err = json.Unmarshal(buf.Bytes(), response); err != nil {
		return fmt.Errorf("decode response body to %T: %w", response, err)
	}

	return nil
}

func prepareQuery(filters ...query.Filter) map[string][]string {
	var queries []string
	for _, f := range filters {
		queries = append(queries, f.Prepare()...)
	}
	return map[string][]string{".query": queries}
}
