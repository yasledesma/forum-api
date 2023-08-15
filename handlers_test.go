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

		if err != nil {
			t.Fatalf("unable to parse response from server: %q into slice %v", response.Body, err)
		}

		got := posts["posts"]
		want := db.Posts
        
		if response.Result().StatusCode != http.StatusOK {
			t.Errorf("expected %v, got %v", http.StatusOK, response.Result().StatusCode)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
