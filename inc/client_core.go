package inc

import (
	"errors"
	"net/http"
)

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {

	client := http.Client{}

	request, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, errors.New("Couldn't make the request")
	}

	for header, value := range headers {
		if len(value) > 0 {
			request.Header.Set(header, value[0])
		}
	}

	response, err := client.Do(request)

	if err != nil {
		return nil, errors.New("Couldn't get the response")
	}

	return response, nil
}
