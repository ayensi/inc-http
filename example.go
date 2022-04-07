package main

import (
	"fmt"
	"github.com/ayensi/inc-http/inc"
	"io/ioutil"
	"net/http"
)

func getGithubClient() inc.HttpHelper {
	client := inc.New()
	commonHeaders := make(http.Header)
	commonHeaders.Set("Accept", "application/xml")
	client.SetHeaders(commonHeaders)
	return client
}

func main() {
	client := getGithubClient()
	response, err := client.GET("https://api.github.com", nil)

	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(response.Body)

	fmt.Println(string(bytes))
}
