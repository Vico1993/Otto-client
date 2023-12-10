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

	articles := client.Chat.ListLatestArticles("1111", "3")
	for _, article := range articles {
		fmt.Println(article.Title)
	}
}
