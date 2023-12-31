package otto

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var baseUrlTest = "http://localhost:8888"

// SetupTest server to facilate test
func setupTest() (client *Client, mux *http.ServeMux, serverURL string, teardown func()) {
	apiHandler := http.NewServeMux()
	apiHandler.Handle(baseUrlTest+"/", http.StripPrefix(baseUrlTest, mux))
	apiHandler.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(os.Stderr, "FAIL: Client.BaseURL path prefix is not preserved in the request URL:")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "\t"+req.URL.String())
		fmt.Fprintln(os.Stderr, "\tDid you accidentally use an absolute endpoint URL rather than relative?")
		http.Error(w, "Client.BaseURL path prefix is not preserved in the request URL.", http.StatusInternalServerError)
	})

	server := httptest.NewServer(apiHandler)
	client = NewClient(nil, server.URL)

	return client, apiHandler, server.URL, server.Close
}

func TestClientGetChatUrl(t *testing.T) {
	client := NewClient(nil, baseUrlTest)

	url := client.GetChatUrl("chatId", "threadId", "/test")
	assert.Equal(t, baseUrlTest+"/chats/chatId/threadId/test", url)
}

func TestClientGetChatUrlWithoutThreadId(t *testing.T) {
	client := NewClient(nil, baseUrlTest)

	url := client.GetChatUrl("chatId", "", "/test")
	assert.Equal(t, baseUrlTest+"/chats/chatId/test", url)
}
