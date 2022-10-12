package headerauthentication

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Config the plugin configuration.
type Config struct {
	Header map[string]string `json:"header,omitempty"`
}

// ErrorResponse represents the response when the API key is invalid .
type ErrorResponse struct {
	ErrorCode string `json:"error_code"`
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
	fmt.Printf("Creating plugin: %s instance: %+v, ctx: %+v\n", name, *config, ctx)
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
		return
	}
	response := ErrorResponse{ErrorCode: "Invalid API Key"}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(403)
	if err := json.NewEncoder(rw).Encode(response); err != nil {
		fmt.Println("Failed to reply request with invalid API Key: " + err.Error())
	}
}
