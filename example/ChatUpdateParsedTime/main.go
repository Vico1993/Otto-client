package main

import (
	"net/http"

	"github.com/Vico1993/Otto-client/otto"
)

var baseUrl = "http://localhost:8888"

func main() {
	client := otto.NewClient(
		&http.Client{},
		baseUrl,
	)

	client.Chat.UpdateParsedTime("1111")
	// fmt.Printf("Chat id: %s", chat.Id)
}
