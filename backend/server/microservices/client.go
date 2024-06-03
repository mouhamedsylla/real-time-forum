// Package microservices defines an HTTP client for interacting with microservices.
package microservices

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// HttpClient represents an HTTP client with a base URL and HTTP method.
type HttpClient struct {
	// BaseURL is the base URL for the service the client will call.
	BaseURL string
	
	// Method is the HTTP method (e.g., GET, POST) that the client will use.
	Method string
}

// NewClient creates a new instance of HttpClient.
func NewClient() *HttpClient {
	return &HttpClient{}
}

// SetBaseURL sets the base URL for the HttpClient.
func (c *HttpClient) SetBaseURL(url string) {
	c.BaseURL = url
}

// SetMethod sets the HTTP method for the HttpClient.
func (c *HttpClient) SetMethod(method string) {
	c.Method = method
}

// Call makes a request to the specified service and endpoint, using the provided request data.
// It decodes the response into the provided response structure and returns an error if any occurs.
func (c *HttpClient) Call(serviceName, endpoint string, request, response interface{}) error {
	// Construct the full URL for the request.
	url := fmt.Sprintf("%s/%s/%s", c.BaseURL, serviceName, endpoint)
	
	// Marshal the request data into JSON.
	rqBody, err := json.Marshal(request)
	if err != nil {
		return err
	}

	// Create a new HTTP request with the specified method, URL, and request body.
	req, err := http.NewRequest(c.Method, url, bytes.NewBuffer(rqBody))
	if err != nil {
		return err
	}
	
	// Set the content type to application/json.
	req.Header.Set("Content-Type", "application/json")
	
	// Create a new HTTP client and send the request.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Decode the response body into the response structure.
	return json.NewDecoder(resp.Body).Decode(response)
}

