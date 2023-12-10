package otto

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type articleCreateBody struct {
	FeedId string   `json:"feed_id"`
	Title  string   `json:"title"`
	Source string   `json:"source"`
	Link   string   `json:"link"`
	Author string   `json:"author"`
	Tags   []string `json:"tags"`
}

func TestCreateArticle(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	feedId := "1"
	title := "Super article"
	source := "source.com"
	link := "https://source.com/article-1"
	author := "John Doe"
	tags := []string{"tags1", "tags2"}

	mux.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		// Assert it's a POST
		assert.Equal(t, r.Method, http.MethodPost)

		body := new(articleCreateBody)
		assert.Nil(t, json.NewDecoder(r.Body).Decode(body))

		assert.Equal(t, feedId, body.FeedId, "Feed Id should match expected value")
		assert.Equal(t, title, body.Title, "title should match expected value")
		assert.Equal(t, source, body.Source, "Source should match expected value")
		assert.Equal(t, link, body.Link, "Link should match expected value")
		assert.Equal(t, author, body.Author, "Author should match expected value")
		assert.Equal(t, tags, body.Tags, "Tags should match expected value")

		article := &Article{
			Id:     "0",
			FeedId: feedId,
			Title:  title,
			Source: source,
			Link:   link,
			Author: author,
			Tags:   tags,
		}

		out, _ := json.Marshal(&ArticleCreateResponse{Article: *article})
		fmt.Fprint(w, string(out))
	})

	res := client.Article.Create(feedId, title, source, link, author, tags)
	assert.Equal(t, res.Title, "Super article")
}

func TestListLatestArticles(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	feedId := "1"
	title := "Super article"
	source := "source.com"
	link := "https://source.com/article-1"
	author := "John Doe"
	tags := []string{"tags1", "tags2"}

	mux.HandleFunc("/chats/chatId/articles/latest", func(w http.ResponseWriter, r *http.Request) {
		// Assert it's a Get
		assert.Equal(t, r.Method, http.MethodGet)

		article := Article{
			Id:     "0",
			FeedId: feedId,
			Title:  title,
			Source: source,
			Link:   link,
			Author: author,
			Tags:   tags,
		}

		out, _ := json.Marshal(&ArticleListLatestResponse{Articles: []Article{article}})
		fmt.Fprint(w, string(out))
	})

	res := client.Chat.ListLatestArticles("chatId")
	assert.Len(t, res, 1)
	assert.Equal(t, res[0].Title, title)
}
