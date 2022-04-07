package inc

import "net/http"

type httpClient struct {
	Headers http.Header
}

type HttpHelper interface {
	SetHeaders(headers http.Header)
	GET(url string, headers http.Header) (*http.Response, error)
	POST(url string, headers http.Header, body interface{}) (*http.Response, error)
	PUT(url string, headers http.Header, body interface{}) (*http.Response, error)
	PATCH(url string, headers http.Header, body interface{}) (*http.Response, error)
	DELETE(url string, headers http.Header) (*http.Response, error)
}

func New() HttpHelper {
	client := &httpClient{}
	return client
}
func (c *httpClient) SetHeaders(headers http.Header) {
	c.Headers = headers
}
func (c *httpClient) GET(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}
func (c *httpClient) POST(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPost, url, headers, body)

}
func (c *httpClient) PUT(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPut, url, headers, body)

}
func (c *httpClient) PATCH(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPatch, url, headers, body)

}
func (c *httpClient) DELETE(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)

}
