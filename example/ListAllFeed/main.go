package main

import (
	"fmt"
	"net/http"

	"github.com/Vico1993/Otto-client/otto"
)

var baseUrl = "http://localhost:8888"

func main() {
	client := otto.NewClient(
		&http.Client{},
		baseUrl,
	)

	feeds := client.Feed.ListAll(false)
	for _, feed := range feeds {
		fmt.Println(feed.Url)
	}
}
