package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetAllPosts(t *testing.T) {
	t.Run("should return all posts", func(t *testing.T) {
		var p []Post
		var posts = map[string][]Post{"posts": p}

		// tests '/api/posts' route
		request, _ := http.NewRequest(http.MethodGet, "", nil)
		response := httptest.NewRecorder()

		handlePosts(response, request)

		err := json.NewDecoder(response.Body).Decode(&posts)
		checkError(t, err)

		got := posts["posts"]
		want := db.Posts

		assertStatusCode(t, response, http.StatusOK)
		assertResponse(t, got, want)
	})
}

func assertResponse(t *testing.T, got, want []Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertStatusCode(t *testing.T, response *httptest.ResponseRecorder, statusCode int) {
	t.Helper()
	if response.Result().StatusCode != statusCode {
		t.Errorf("expected %v, got %v", http.StatusOK, response.Result().StatusCode)
	}
}

func checkError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}
