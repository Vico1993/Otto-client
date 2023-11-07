package otto

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Chat struct {
	Id               string     `json:"Id"`
	TelegramChatId   string     `json:"TelegramChatId"`
	TelegramUserId   string     `json:"TelegramUserId"`
	TelegramThreadId string     `json:"TelegramThreadId"`
	Tags             []string   `json:"Tags"`
	LastTimeParsed   *time.Time `json:"LastTimeParsed"`
}

type ChatCreateResponse struct {
	Chat Chat `json:"chat"`
}

type ChatService service

// Create the chat in the Otto API
func (s *ChatService) Create(chatId string, userId string, threadId string, tags []string) *Chat {
	dataStr := `{
		"chat_id": "` + chatId + `",
		"user_id": "` + userId + `",
		"thread_id": "` + threadId + `",
		"tags": []
	}`

	if threadId == "" {
		dataStr = strings.Replace(dataStr, `"thread_id": "",`, "", 1)
	}

	data := []byte(dataStr)
	req, err := http.NewRequest(
		http.MethodPost,
		s.client.BaseURL+"/chats", // TODO: Change ? move this in main.go
		strings.NewReader(
			string(data),
		),
	)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		fmt.Println("Error creating the request to initiate chat: " + err.Error())
		return nil
	}

	body, err := s.client.Do(req)
	if err != nil {
		return nil
	}

	var res ChatCreateResponse
	_ = json.Unmarshal(body, &res)

	return &res.Chat
}
