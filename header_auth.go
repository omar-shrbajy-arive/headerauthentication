package headerauthentication

import (
	"context"
	"fmt"
	"net/http"
)

// Config the plugin configuration.
type Config struct {
	Header map[string]string `json:"header,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Header: map[string]string{
			"name": "X-API-KEY",
		},
	}
}

type HeaderAuth struct {
	next   http.Handler
	header map[string]string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if len(config.Header) == 0 {
		return nil, fmt.Errorf("headers cannot be empty")
	}

	return &HeaderAuth{
		header: config.Header,
		next:   next,
	}, nil
}

func (a *HeaderAuth) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.Header.Get(a.header["name"]) == a.header["key"] {
		req.Header.Del(a.header["name"])
		a.next.ServeHTTP(rw, req)
	}
	http.Error(rw, "Not allowed - verified null", http.StatusUnauthorized)
}
