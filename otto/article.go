package otto

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Article struct {
	Id     string   `json:"Id"`
	Title  string   `json:"title"`
	Source string   `json:"source"`
	Author string   `json:"author"`
	Link   string   `json:"link"`
	Tags   []string `json:"tags"`
	FeedId string   `json:"feed_id"`
}

type ArticleCreateResponse struct {
	Article Article `json:"article"`
}

type ArticleListLatestResponse struct {
	Articles []Article `json:"articles"`
}

type ArticleService service

// Create a Article for a specific feed Id
func (s *ArticleService) Create(feedId string, title string, source string, link string, author string, tags []string) *Article {
	data := []byte(`{
		"feed_id": "` + feedId + `",
		"title": "` + title + `",
		"source": "` + source + `",
		"link": "` + link + `",
		"author": "` + author + `",
		"tags": ["` + strings.Join(tags, "\",\"") + `"]
	}`)

	req, err := http.NewRequest(
		http.MethodPost,
		s.client.BaseURL+"/articles",
		strings.NewReader(
			string(data),
		),
	)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		fmt.Println("Error creating the request to add article: " + err.Error())
		return nil
	}

	body, err := s.client.Do(req)
	if err != nil {
		fmt.Println("Error requesting new article: " + err.Error())
		return nil
	}

	var res ArticleCreateResponse
	_ = json.Unmarshal(body, &res)

	return &res.Article
}

// ListLatestArticles all latest articles from chats
func (s *ChatService) ListLatestArticles(chatId string) []Article {
	req, err := http.NewRequest(
		http.MethodGet,
		s.client.BaseURL+"/chats/"+chatId+"/articles/latest",
		strings.NewReader(
			string([]byte{}),
		),
	)
	if err != nil {
		fmt.Println("Error creating the request to list all latest articles from chat: " + err.Error())
		return nil
	}

	body, err := s.client.Do(req)
	if err != nil {
		return nil
	}

	var res ArticleListLatestResponse
	_ = json.Unmarshal(body, &res)

	return res.Articles
}
