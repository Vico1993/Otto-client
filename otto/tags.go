package otto

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type TagsListResponse struct {
	Tags []string `json:"tags"`
}

type TagsCreateResponse struct {
	Tags []string `json:"tags"`
}

type TagsDeleteResponse struct {
	Deleted bool `json:"deleted"`
}

type TagService service

// List all chats from a chat id and thread id
func (s *TagService) List(chatId string, threadId string) []string {
	req, err := http.NewRequest(
		http.MethodGet,
		s.client.GetChatUrl(chatId, threadId, "/tags"),
		strings.NewReader(
			string([]byte{}),
		),
	)
	if err != nil {
		fmt.Println("Error creating the request to list tags: " + err.Error())
		return nil
	}

	body, err := s.client.Do(req)
	if err != nil {
		return nil
	}

	var res TagsListResponse
	_ = json.Unmarshal(body, &res)

	return res.Tags
}

// Delete tag from a chat id and thread id
func (s *TagService) Delete(chatId string, threadId string, tag string) bool {
	req, err := http.NewRequest("DELETE", s.client.GetChatUrl(chatId, threadId, "/tags/"+tag), nil)
	if err != nil {
		fmt.Println("Error creating the request to delete tag in chat: " + err.Error())
		return false
	}

	body, err := s.client.Do(req)
	if err != nil {
		fmt.Println("Error deleting tag: " + err.Error())
		return false
	}

	var res TagsDeleteResponse
	_ = json.Unmarshal(body, &res)

	return res.Deleted
}

// Add tag from a chat id and thread id
func (s *TagService) Create(chatId string, threadId string, tags []string) []string {
	data := []byte(`{
		"tags": ["` + strings.Join(tags, "\",\"") + `"]
	}`)

	req, err := http.NewRequest(
		http.MethodPost,
		s.client.GetChatUrl(chatId, threadId, "/tags"),
		strings.NewReader(
			string(data),
		),
	)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		fmt.Println("Error creating the request to add tag: " + err.Error())
		return nil
	}

	body, err := s.client.Do(req)
	if err != nil {
		fmt.Println("Error requesting new tag: " + err.Error())
		return nil
	}

	var res TagsCreateResponse
	_ = json.Unmarshal(body, &res)

	return res.Tags
}
