package go_routeros_client

import (
	"crypto/tls"
	"net/http"
	"time"
)

// Option allows to customize the Client
type Option func(c *Client)

// WithTransport sets a custom http.RoundTripper transport for the HTTP client
func WithTransport(rt http.RoundTripper) Option {
	return func(c *Client) {
		c.httpClient.Transport = rt
	}
}

// WithHTTPTimeout sets a custom HTTP client timeout
func WithHTTPTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

// WithInsecureTLS enables the HTTP client to establish insecure TLS connections
func WithInsecureTLS() Option {
	return func(c *Client) {
		tr, ok := c.httpClient.Transport.(*http.Transport)
		if !ok {
			return
		}

		if tr.TLSClientConfig == nil {
			tr.TLSClientConfig = &tls.Config{
				MinVersion: tls.VersionTLS12,
			}
		}

		tr.TLSClientConfig.InsecureSkipVerify = true
	}
}
