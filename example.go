package main

import (
	"encoding/json"
	"fmt"
	"github.com/ayensi/inc-http/inc"
	"io/ioutil"
	"net/http"
)

var (
	githubClient = getGithubClient()
)

func getGithubClient() inc.HttpHelper {
	client := inc.New()
	commonHeaders := make(http.Header)
	commonHeaders.Set("Accept", "application/json")
	client.SetHeaders(commonHeaders)
	return client
}

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	user := new(User)
	user.FirstName = "Okan"
	user.LastName = "Altun"
	createUser(user)
}
func getUrls() {

	response, err := githubClient.GET("https://api.github.com", nil)

	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(response.Body)

	fmt.Println(string(bytes))
}
func createUser(user *User) {
	requestBody, _ := json.Marshal(user)
	response, err := githubClient.POST("https://api.github.com", nil, requestBody)

	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(response.Body)

	fmt.Println(string(bytes))
}
