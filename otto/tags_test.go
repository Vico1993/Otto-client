package otto

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListTags(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/chats/chatId/tags", func(w http.ResponseWriter, r *http.Request) {
		// Assert it's a GET
		assert.Equal(t, r.Method, http.MethodGet)

		out, _ := json.Marshal(&TagsListResponse{Tags: []string{"test", "test2"}})
		fmt.Fprint(w, string(out))
	})

	res := client.Tag.List("chatId", "")

	assert.Len(t, res, 2)
	assert.Equal(t, res, []string{"test", "test2"})
}

func TestCreateTag(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/chats/chatId/tags", func(w http.ResponseWriter, r *http.Request) {
		// Assert it's a POST
		assert.Equal(t, r.Method, http.MethodPost)

		out, _ := json.Marshal(&TagsCreateResponse{Tags: []string{"test"}})
		fmt.Fprint(w, string(out))
	})

	res := client.Tag.Create("chatId", "", []string{"test"})
	assert.Equal(t, res, []string{"test"})
}

func TestDeleteTag(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/chats/chatId/tags/test", func(w http.ResponseWriter, r *http.Request) {
		// Assert it's a DELETE
		assert.Equal(t, r.Method, http.MethodDelete)

		out, _ := json.Marshal(&TagsDeleteResponse{Deleted: true})
		fmt.Fprint(w, string(out))
	})

	res := client.Tag.Delete("chatId", "", "test")
	assert.Equal(t, res, true)
}
