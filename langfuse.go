package langfuse

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"

	"github.com/google/go-querystring/query"
)

const (
	// BaseURLCloud is the EU region endpoint (default)
	BaseURLCloud = "https://cloud.langfuse.com"
	// BaseURLCloudUS is the US region endpoint
	BaseURLCloudUS = "https://us.cloud.langfuse.com"
	// BaseURLCloudJP is the Japan region endpoint
	BaseURLCloudJP = "https://jp.cloud.langfuse.com"
)

var defaultBaseURL, _ = url.Parse(BaseURLCloud)

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Client is the Langfuse API client.
type Client struct {
	publicKey  string
	secretKey  string
	baseURL    *url.URL
	debug      bool
	log        ilogger
	httpclient httpClient
}

// Option defines a functional option for Client.
type Option func(*Client)

// OptionBaseURL sets a custom base URL.
func OptionBaseURL(endpoint string) func(*Client) {
	baseURL, _ := url.Parse(endpoint)
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// OptionHTTPClient sets a custom HTTP client.
func OptionHTTPClient(client httpClient) func(*Client) {
	return func(c *Client) {
		c.httpclient = client
	}
}

// OptionDebug enables debug logging.
func OptionDebug(b bool) func(*Client) {
	return func(c *Client) {
		c.debug = b
	}
}

// OptionLog sets a custom logger.
func OptionLog(l logger) func(*Client) {
	return func(c *Client) {
		c.log = internalLog{logger: l}
	}
}

// New creates a new Langfuse client with the given public/secret key pair.
func New(publicKey, secretKey string, options ...Option) *Client {
	c := &Client{
		publicKey:  publicKey,
		secretKey:  secretKey,
		baseURL:    defaultBaseURL,
		httpclient: &http.Client{},
		log:        log.New(os.Stderr, "kenzo0107/langfuse", log.LstdFlags|log.Lshortfile),
	}

	for _, opt := range options {
		opt(c)
	}

	return c
}

// NewFromEnv creates a client from LANGFUSE_PUBLIC_KEY, LANGFUSE_SECRET_KEY, and LANGFUSE_HOST environment variables.
func NewFromEnv(options ...Option) *Client {
	publicKey := os.Getenv("LANGFUSE_PUBLIC_KEY")
	secretKey := os.Getenv("LANGFUSE_SECRET_KEY")

	opts := options
	if host := os.Getenv("LANGFUSE_HOST"); host != "" {
		opts = append([]Option{OptionBaseURL(host)}, opts...)
	}

	return New(publicKey, secretKey, opts...)
}

// Debugf prints a formatted debug message.
func (c *Client) Debugf(format string, v ...interface{}) {
	if c.debug {
		if err := c.log.Output(2, fmt.Sprintf(format, v...)); err != nil {
			c.Debugln(err)
		}
	}
}

// Debugln prints a debug message.
func (c *Client) Debugln(v ...interface{}) {
	if c.debug {
		if err := c.log.Output(2, fmt.Sprintln(v...)); err != nil {
			_ = c.log.Output(2, fmt.Sprintf("debug log error: %v", err))
		}
	}
}

// Debug returns true if debug mode is enabled.
func (c *Client) Debug() bool {
	return c.debug
}

// Bool returns a pointer to the given bool value.
func Bool(v bool) *bool { return &v }

// String returns a pointer to the given string value.
func String(v string) *string { return &v }

// Int returns a pointer to the given int value.
func Int(v int) *int { return &v }

// Float64 returns a pointer to the given float64 value.
func Float64(v float64) *float64 { return &v }

// NewRequest creates an HTTP request with Basic Auth.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if c.baseURL == nil {
		return nil, fmt.Errorf("baseURL is nil")
	}

	u, err := c.baseURL.Parse(c.baseURL.Path + urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		if er := enc.Encode(body); er != nil {
			return nil, er
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.publicKey, c.secretKey)

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

// Do sends the request and decodes the response into v.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) error {
	if ctx == nil {
		return errors.New("context must be non-nil")
	}

	req = req.WithContext(ctx)

	resp, err := c.httpclient.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		return err
	}
	defer func() {
		if er := resp.Body.Close(); er != nil {
			err = er
		}
	}()

	if err = checkStatusCode(resp, c); err != nil {
		return err
	}

	if err = decodeResponse(resp.Body, v); err != nil {
		return err
	}

	return err
}

func decodeResponse(body io.Reader, v interface{}) error {
	if v == nil {
		return nil
	}

	if w, ok := v.(io.Writer); ok {
		_, err := io.Copy(w, body)
		return err
	}

	decErr := json.NewDecoder(body).Decode(v)
	if errors.Is(decErr, io.EOF) {
		return nil
	}

	return decErr
}

// AddOptions appends struct fields as URL query parameters.
func (c *Client) AddOptions(s string, opts interface{}) (string, error) {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opts)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()

	return u.String(), nil
}

// Pagination holds paging metadata returned by list endpoints.
type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalItems int `json:"totalItems"`
	TotalPages int `json:"totalPages"`
}
