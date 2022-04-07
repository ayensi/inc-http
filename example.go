package main

import (
	"fmt"
	"github.com/ayensi/inc-http/inc"
	"io/ioutil"
)

func main() {
	client := inc.New()
	response, err := client.GET("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}
