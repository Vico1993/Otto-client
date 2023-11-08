package otto

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListFeed(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/chats/chatId/feeds", func(w http.ResponseWriter, r *http.Request) {
		// Assert it's a GET
		assert.Equal(t, r.Method, http.MethodGet)

		feed := &Feed{
			Id:  "1",
			Url: "https://google.com",
		}

		out, _ := json.Marshal(&FeedListResponse{Feeds: []Feed{*feed}})
		fmt.Fprint(w, string(out))
	})

	res := client.Feed.List("chatId", "")

	assert.Len(t, res, 1)
	assert.Equal(t, res[0].Url, "https://google.com")
}

func TestCreateFeed(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/feeds", func(w http.ResponseWriter, r *http.Request) {
		// Assert it's a POST
		assert.Equal(t, r.Method, http.MethodPost)

		feed := &Feed{
			Id:  "1",
			Url: "https://google.com",
		}

		out, _ := json.Marshal(&FeedCreateResponse{Feed: *feed})
		fmt.Fprint(w, string(out))
	})

	res := client.Feed.Create("https://google.com")
	assert.Equal(t, res.Url, "https://google.com")
}

func TestUnlinkFeed(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/chats/chatId/feeds/feedId", func(w http.ResponseWriter, r *http.Request) {
		// Assert it's a DELETE
		assert.Equal(t, r.Method, http.MethodDelete)

		fmt.Fprint(w, "{\"deleted\": true}")
	})

	res := client.Feed.UnLink("chatId", "", "feedId")
	assert.True(t, res)
}

func TestLinkFeed(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/chats/chatId/feeds/feedId", func(w http.ResponseWriter, r *http.Request) {
		// Assert it's a POST
		assert.Equal(t, r.Method, http.MethodPost)

		fmt.Fprint(w, "{\"added\": true}")
	})

	res := client.Feed.Link("chatId", "", "feedId")
	assert.True(t, res)
}
