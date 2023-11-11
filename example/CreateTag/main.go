package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Vico1993/Otto-client/otto"
)

var baseUrl = "http://localhost:8888"

func main() {
	client := otto.NewClient(
		&http.Client{},
		baseUrl,
	)

	tags := client.Tag.Create("1111", "3", []string{"test"})
	fmt.Println("tags: " + strings.Join(tags, ", "))
}
