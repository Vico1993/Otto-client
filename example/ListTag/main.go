package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Vico1993/Otto-client/v1/otto"
)

var baseUrl = "http://localhost:8888"

func main() {
	client := otto.NewClient(
		&http.Client{},
		baseUrl,
	)

	tags := client.Tag.List("1111", "3")
	fmt.Println("tags: " + strings.Join(tags, ", "))
}
