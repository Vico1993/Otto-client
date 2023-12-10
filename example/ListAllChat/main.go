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

	chats := client.Chat.ListAll()
	for _, chat := range chats {
		fmt.Println(chat.TelegramChatId)
	}
}
