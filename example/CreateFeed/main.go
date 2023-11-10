package main

import (
	"fmt"
	"net/http"

	"github.com/Vico1993/Otto-client/v1/otto"
)

var baseUrl = "http://localhost:8888"

func main() {
	client := otto.NewClient(
		&http.Client{},
		baseUrl,
	)

	feed := client.Feed.Create("https://google.com")
	fmt.Println(feed.Id)
}
