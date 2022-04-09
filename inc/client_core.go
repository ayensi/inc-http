package inc

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
)

func (c *httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	switch contentType {
	case "application/json":
		return json.Marshal(body)
	case "application/xml":
		return xml.Marshal(body)
	default:
		return json.Marshal(body)
	}
}

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {

	client := http.Client{}
	fullHeaders := c.getRequestHeaders(headers)
	requestBody, err := c.getRequestBody(fullHeaders.Get("Content-Type"), body)
	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, errors.New("Couldn't make the request")
	}

	request.Header = fullHeaders
	response, err := client.Do(request)

	if err != nil {
		return nil, errors.New("Couldn't get the response")
	}

	return response, nil
}
func (c *httpClient) getRequestHeaders(headers http.Header) http.Header {
	result := make(http.Header)
	for header, value := range c.Headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	for header, value := range headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}
	return result
}
