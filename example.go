package main

import (
	"fmt"
	"github.com/ayensi/inc-http/inc"
	"io/ioutil"
	"net/http"
)

func main() {
	client := inc.New()
	headers := make(http.Header)
	headers.Set("Authorization", "Bearer ABC-123")
	response, err := client.GET("https://api.github.com", headers)

	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(response.Body)

	fmt.Println(string(bytes))
}
