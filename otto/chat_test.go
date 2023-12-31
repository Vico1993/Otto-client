package otto

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type chatCreateBody struct {
	ChatId   string   `json:"chat_id"`
	UserId   string   `json:"user_id"`
	ThreadId string   `json:"thread_id,omitempty"`
	Tags     []string `json:"tags"`
}

func TestChatCreate(t *testing.T) {
	chatId := "chatId"
	userId := "userId"
	threadId := "threadId"
	tags := []string{}

	expectedResponse := Chat{
		Id:               "1",
		TelegramChatId:   chatId,
		TelegramUserId:   userId,
		TelegramThreadId: threadId,
		Tags:             tags,
		LastTimeParsed:   nil,
	}

	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/chats", func(w http.ResponseWriter, r *http.Request) {
		// Assert it's a post
		assert.Equal(t, r.Method, http.MethodPost)

		body := new(chatCreateBody)
		assert.Nil(t, json.NewDecoder(r.Body).Decode(body))

		assert.Equal(t, chatId, body.ChatId, "Chat Id should match expected value")
		assert.Equal(t, userId, body.UserId, "User Id should match expected value")
		assert.Equal(t, threadId, body.ThreadId, "Thread Id should match expected value")
		assert.Equal(t, tags, body.Tags, "Tags should match expected value")

		out, _ := json.Marshal(&ChatCreateResponse{Chat: expectedResponse})
		fmt.Fprint(w, string(out))
	})

	res := client.Chat.Create(chatId, userId, threadId, tags)

	assert.Equal(t, expectedResponse, *res, "Chat response must equal to chat expected")
}

func TestChatCreateWithoutThreadId(t *testing.T) {
	chatId := "chatId"
	userId := "userId"
	tags := []string{}

	expectedResponse := Chat{
		Id:               "1",
		TelegramChatId:   chatId,
		TelegramUserId:   userId,
		TelegramThreadId: "",
		Tags:             tags,
		LastTimeParsed:   nil,
	}

	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/chats", func(w http.ResponseWriter, r *http.Request) {
		// Assert it's a post
		assert.Equal(t, r.Method, http.MethodPost)

		body := new(chatCreateBody)
		assert.Nil(t, json.NewDecoder(r.Body).Decode(body))

		assert.Equal(t, chatId, body.ChatId, "Chat Id should match expected value")
		assert.Equal(t, userId, body.UserId, "User Id should match expected value")
		assert.Equal(t, "", body.ThreadId, "Thread Id should be empty")
		assert.Equal(t, tags, body.Tags, "Tags should match expected value")

		out, _ := json.Marshal(&ChatCreateResponse{Chat: expectedResponse})
		fmt.Fprint(w, string(out))
	})

	res := client.Chat.Create(chatId, userId, "", tags)

	assert.Equal(t, expectedResponse, *res, "Chat response must equal to chat expected")
}

func TestChatGetAll(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	chatId := "chatId"
	userId := "userId"
	threadId := "threadId"
	tags := []string{}

	expectedResponse := Chat{
		Id:               "1",
		TelegramChatId:   chatId,
		TelegramUserId:   userId,
		TelegramThreadId: threadId,
		Tags:             tags,
		LastTimeParsed:   nil,
	}

	mux.HandleFunc("/chats", func(w http.ResponseWriter, r *http.Request) {
		// Assert it's a GET
		assert.Equal(t, r.Method, http.MethodGet)

		out, _ := json.Marshal(&ChatGetAllResponse{Chats: []Chat{expectedResponse}})
		fmt.Fprint(w, string(out))
	})

	res := client.Chat.ListAll()

	assert.Len(t, res, 1)
	assert.Equal(t, res[0].TelegramChatId, chatId)
}

func TestUpdateParsedTime(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/chats/chatId/threadId/parsed", func(w http.ResponseWriter, r *http.Request) {
		// Assert it's a GET
		assert.Equal(t, r.Method, http.MethodGet)

		w.WriteHeader(http.StatusNoContent)
	})

	client.Chat.UpdateParsedTime("chatId", "threadId")
}

func TestUpdateParsedTimeWithoutThreadId(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	mux.HandleFunc("/chats/chatId/parsed", func(w http.ResponseWriter, r *http.Request) {
		// Assert it's a GET
		assert.Equal(t, r.Method, http.MethodGet)

		w.WriteHeader(http.StatusNoContent)
	})

	client.Chat.UpdateParsedTime("chatId", "")
}
