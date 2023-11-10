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

	done := client.Tag.Delete("1111", "3", "test")
	fmt.Println("Is it done? ", done)
}
