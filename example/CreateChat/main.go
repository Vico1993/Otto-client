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

	chat := client.Chat.Create("1111", "22", "3", []string{})
	fmt.Printf("Chat id: %s", chat.Id)
}
