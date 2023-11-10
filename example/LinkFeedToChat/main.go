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

	done := client.Feed.Link("1111", "3", "ffdc9945-eeea-42e5-a293-83455fb3d0c5")
	fmt.Println("Is it done? ", done)
}
