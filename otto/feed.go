package otto

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Feed struct {
	Id  string `json:"Id"`
	Url string `json:"Url"`
}

type FeedListResponse struct {
	Feeds []Feed `json:"feeds"`
}

type FeedListArticlesResponse struct {
	Articles []Article `json:"articles"`
}

type FeedCreateResponse struct {
	Feed Feed `json:"feed"`
}

type FeedDisabledResponse struct {
	Deleted bool `json:"deleted"`
}

type FeedLinkResponse struct {
	Added bool `json:"added"`
}

type FeedService service

// List all feeds from a chat id and thread id
func (s *FeedService) List(chatId string, threadId string) []Feed {
	req, err := http.NewRequest(
		http.MethodGet,
		s.client.GetChatUrl(chatId, threadId, "/feeds"),
		strings.NewReader(
			string([]byte{}),
		),
	)
	if err != nil {
		fmt.Println("Error creating the request to initiate chat: " + err.Error())
		return nil
	}

	body, err := s.client.Do(req)
	if err != nil {
		return nil
	}

	var res FeedListResponse
	_ = json.Unmarshal(body, &res)

	return res.Feeds
}

// Create a feed for Otto
func (s *FeedService) Create(url string) *Feed {
	data := []byte(`{
		"url": "` + url + `"
	}`)

	req, err := http.NewRequest(
		http.MethodPost,
		s.client.BaseURL+"/feeds",
		strings.NewReader(
			string(data),
		),
	)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		fmt.Println("Error creating the request to add feed: " + err.Error())
		return nil
	}

	body, err := s.client.Do(req)
	if err != nil {
		fmt.Println("Error requesting new feed: " + err.Error())
		return nil
	}

	var res FeedCreateResponse
	_ = json.Unmarshal(body, &res)

	return &res.Feed
}

// UnLink a feed from a chat id and thread id
func (s *FeedService) UnLink(chatId string, threadId string, feedId string) bool {
	req, err := http.NewRequest("DELETE", s.client.GetChatUrl(chatId, threadId, "/feeds/"+feedId), nil)
	if err != nil {
		fmt.Println("Error creating the request to disabled the feed in chat: " + err.Error())
		return false
	}

	body, err := s.client.Do(req)
	if err != nil {
		fmt.Println("Error deleting tag: " + err.Error())
		return false
	}

	var res FeedDisabledResponse
	_ = json.Unmarshal(body, &res)

	return res.Deleted
}

// Link a feed from a chat id and thread id
func (s *FeedService) Link(chatId string, threadId string, feedId string) bool {
	req, err := http.NewRequest(
		http.MethodPost,
		s.client.GetChatUrl(chatId, threadId, "/feeds/"+feedId),
		strings.NewReader(
			string([]byte{}),
		),
	)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		fmt.Println("Error linking chat and feed: " + err.Error())
		return false
	}

	body, err := s.client.Do(req)
	if err != nil {
		fmt.Println("Error linking chat and feed: " + err.Error())
		return false
	}

	var res FeedLinkResponse
	_ = json.Unmarshal(body, &res)

	return res.Added
}

// List all Feeds
// With active parameters you can retrieve only active feeds
func (s *FeedService) ListAll(active bool) []Feed {
	url := s.client.BaseURL + "/feeds"
	if active {
		url = url + "/active"
	}

	req, err := http.NewRequest(
		http.MethodGet,
		url,
		strings.NewReader(
			string([]byte{}),
		),
	)
	if err != nil {
		fmt.Println("Error creating the request to list all feeds: " + err.Error())
		return nil
	}

	body, err := s.client.Do(req)
	if err != nil {
		return nil
	}

	var res FeedListResponse
	_ = json.Unmarshal(body, &res)

	return res.Feeds
}

// Retrieve Feed articles
func (s *FeedService) ListArticles(feedId string) []Article {
	req, err := http.NewRequest(
		http.MethodGet,
		s.client.BaseURL+"/feeds/"+feedId+"/articles",
		strings.NewReader(
			string([]byte{}),
		),
	)
	if err != nil {
		fmt.Println("Error creating the request to list articles: " + err.Error())
		return nil
	}

	body, err := s.client.Do(req)
	if err != nil {
		return nil
	}

	var res FeedListArticlesResponse
	_ = json.Unmarshal(body, &res)

	return res.Articles
}
