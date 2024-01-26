package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestGetAllPosts(t *testing.T) {
	t.Run("should return all posts", func(t *testing.T) {
		var p []Post
		var posts = map[string][]Post{"posts": p}

		// tests '/api/posts' route
		request, _ := http.NewRequest(http.MethodGet, "/api/posts/", nil)
		response := httptest.NewRecorder()

		GetPosts(response, request, nil)

		err := json.NewDecoder(response.Body).Decode(&posts)
		checkError(t, err)

		got := posts["posts"]
		want := db.Posts

		assertStatusCode(t, response, http.StatusOK)
		assertResponse(t, got, want)
	})
}

func TestCreatePost(t *testing.T) {
	t.Run("should succeed on post requests", func(t *testing.T) {
		var p Post

		// tests '/api/posts' route
		request, _ := http.NewRequest(http.MethodPost, "/api/posts/", strings.NewReader(`{"title":"Bobby","link":"https://beautyful-pictures.webs.com/photos/Cute-Puppies/puppy-lab-HD-wallpaper.jpg","username":"ILoveDogs96"}`))
		response := httptest.NewRecorder()
		request.Header.Add("Content-Type", "application/json")

		AddPost(response, request, nil)

		err := json.NewDecoder(response.Body).Decode(&p)
		checkError(t, err)

		got := p
		want := db.Posts[len(db.Posts)-1]

		assertStatusCode(t, response, http.StatusCreated)

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("should fail if body has unknown property", func(t *testing.T) {
		var error map[string]string

		// tests '/api/posts' route
		request, _ := http.NewRequest(http.MethodPost, "/api/posts/", strings.NewReader(`{"title":"Bobby","link":"https://beautyful-pictures.webs.com/photos/Cute-Puppies/puppy-lab-HD-wallpaper.jpg","user":"ILoveDogs96"}`))
		response := httptest.NewRecorder()
		request.Header.Add("Content-Type", "application/json")

		AddPost(response, request, nil)

		err := json.NewDecoder(response.Body).Decode(&error)
		checkError(t, err)

		got := error
		want := map[string]string{"error": "couldn't add post"}

		assertStatusCode(t, response, http.StatusBadRequest)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestGetSinglePost(t *testing.T) {
	t.Run("should successfully return a single post", func(t *testing.T) {
		var p Post
		var id = 2
        var pid = httprouter.Param{Key: "pid", Value: fmt.Sprint(id)}

		// tests '/api/posts/:id' route
		request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/posts/%d/", id), nil)
		response := httptest.NewRecorder()

		GetPost(response, request, []httprouter.Param{pid})

		err := json.NewDecoder(response.Body).Decode(&p)
		checkError(t, err)

		got := p
		want := db.Posts[id-1]

		assertStatusCode(t, response, http.StatusOK)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("should fail with non-existent post id", func(t *testing.T) {
		var error map[string]string
		var id = 483
        var pid = httprouter.Param{Key: "pid", Value: fmt.Sprint(id)}
        
		// tests '/api/posts/:id' route
		request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/posts/%d/", id), nil)
		response := httptest.NewRecorder()

		GetPost(response, request, []httprouter.Param{pid})

		err := json.NewDecoder(response.Body).Decode(&error)
		checkError(t, err)

		got := error
		want := map[string]string{"error": "post not found"}
        
		assertStatusCode(t, response, http.StatusNotFound)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestGetAllCommentsForPost(t *testing.T) {
	t.Run("should return all comments for post", func(t *testing.T) {
		var c []Comment
		var comments = map[string][]Comment{"comments": c}
		var id = 1
        var pid = httprouter.Param{Key: "pid", Value: fmt.Sprint(id)}

		// tests '/api/posts/:id/comments' route
		request, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/api/posts/%d/comments/", id), nil)
		response := httptest.NewRecorder()

		GetComments(response, request, []httprouter.Param{pid})

		err := json.NewDecoder(response.Body).Decode(&comments)
		checkError(t, err)

		got := comments["comments"][id - 1].PostId
		want := db.Comments[id - 1].PostId

		assertStatusCode(t, response, http.StatusOK)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})


	t.Run("should fail with non-existent post id", func(t *testing.T) {
		var error map[string]string
		var id = 483
        var pid = httprouter.Param{Key: "pid", Value: fmt.Sprint(id)}

		// tests '/api/posts/:id/comments' route
		request, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/api/posts/%d/comments/", id), nil)
		response := httptest.NewRecorder()

		GetComments(response, request, []httprouter.Param{pid})

		err := json.NewDecoder(response.Body).Decode(&error)
		checkError(t, err)

		got := error
		want := map[string]string{"error": "comments not found"}

		assertStatusCode(t, response, http.StatusNotFound)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})


	t.Run("should fail for posts with no comments", func(t *testing.T) {
		var error map[string]string
		var id = 2
        var pid = httprouter.Param{Key: "pid", Value: fmt.Sprint(id)}

		// tests '/api/posts/:id/comments' route
		request, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/api/posts/%d/comments/", id), nil)
		response := httptest.NewRecorder()

		GetComments(response, request, []httprouter.Param{pid})

		err := json.NewDecoder(response.Body).Decode(&error)
		checkError(t, err)

		got := error
		want := map[string]string{"error": "comments not found"}

		assertStatusCode(t, response, http.StatusNotFound)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestUpdateComment(t *testing.T) {
    var body = `{"text":"so baby! edit: change my mind, it's a cat"}`
    
	t.Run("should update existing comment", func(t *testing.T) {
		var c Comment
        var postId = 1
        var commentId = 1
        var pid = httprouter.Param{Key: "pid", Value: fmt.Sprint(postId)}
        var cid = httprouter.Param{Key: "cid", Value: fmt.Sprint(commentId)}

        // tests '/api/posts/:pid/comments/:cid' route
		request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/api/posts/%d/comments/%d/", postId, commentId), strings.NewReader(body))
		response := httptest.NewRecorder()
		request.Header.Add("Content-Type", "application/json")

		UpdateComment(response, request, []httprouter.Param{pid, cid})

		err := json.NewDecoder(response.Body).Decode(&c)
		checkError(t, err)

		got := c.Text
		want := db.Comments[c.Id - 1].Text

		assertStatusCode(t, response, http.StatusOK)

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("should fail if comment doesn't exist", func(t *testing.T) {
		var error map[string]string
        var postId = 1
        var commentId = 234
        var pid = httprouter.Param{Key: "pid", Value: fmt.Sprint(postId)}
        var cid = httprouter.Param{Key: "cid", Value: fmt.Sprint(commentId)}

        // tests '/api/posts/:postId/comments/:commentId' route
		request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/api/posts/%d/comments/%d/", postId, commentId), strings.NewReader(body))
		response := httptest.NewRecorder()
		request.Header.Add("Content-Type", "application/json")

		UpdateComment(response, request, []httprouter.Param{pid, cid})

		err := json.NewDecoder(response.Body).Decode(&error)
		checkError(t, err)

		got := error
		want := map[string]string{"error": "comment not found"}

		assertStatusCode(t, response, http.StatusNotFound)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("should fail if post doesn't exist", func(t *testing.T) {
		var error map[string]string
        var postId = 234
        var commentId = 2
        var pid = httprouter.Param{Key: "pid", Value: fmt.Sprint(postId)}
        var cid = httprouter.Param{Key: "cid", Value: fmt.Sprint(commentId)}

        // tests '/api/posts/:postId/comments/:commentId' route
		request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/api/posts/%d/comments/%d/", postId, commentId), strings.NewReader(body))
		response := httptest.NewRecorder()
		request.Header.Add("Content-Type", "application/json")

		UpdateComment(response, request, []httprouter.Param{pid, cid})

		err := json.NewDecoder(response.Body).Decode(&error)
		checkError(t, err)

		got := error
		want := map[string]string{"error": "comment not found"}

		assertStatusCode(t, response, http.StatusNotFound)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestDeleteSinglePost(t *testing.T) {
	t.Run("should successfully delete a single post", func(t *testing.T) {
		var p Post
		var id = 1
        var pid = httprouter.Param{Key: "pid", Value: fmt.Sprint(id)}

		// tests '/api/posts/:id' route
		request, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/api/posts/%d/", id), nil)
		response := httptest.NewRecorder()

		DeletePost(response, request, []httprouter.Param{pid})

		err := json.NewDecoder(response.Body).Decode(&p)
		checkError(t, err)

		got := p.Id
		want := id

		assertStatusCode(t, response, http.StatusOK)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("should fail with non-existent post id", func(t *testing.T) {
		var error map[string]string
		var id = 483
        var pid = httprouter.Param{Key: "pid", Value: fmt.Sprint(id)}

		// tests '/api/posts/:id' route
		request, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/api/posts/%d/", id), nil)
		response := httptest.NewRecorder()

		DeletePost(response, request, []httprouter.Param{pid})

		err := json.NewDecoder(response.Body).Decode(&error)
		checkError(t, err)

		got := error
		want := map[string]string{"error": "can't delete post"}

		assertStatusCode(t, response, http.StatusNotFound)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestCreateComment(t *testing.T) {
	t.Run("should succeed on adding a comment", func(t *testing.T) {
		var c Comment
        var id = 1
        var pid = httprouter.Param{Key: "pid", Value: fmt.Sprint(id)}

		// tests '/api/posts/:id/comments' route
		request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/api/posts/%d/comments/", id), strings.NewReader(`{"text":"so baby!","username":"ILoveDogs96"}`))
		response := httptest.NewRecorder()
		request.Header.Add("Content-Type", "application/json")

		AddComment(response, request, []httprouter.Param{pid})

		err := json.NewDecoder(response.Body).Decode(&c)
		checkError(t, err)

		got := c
		want := db.Comments[c.Id - 1]

		assertStatusCode(t, response, http.StatusCreated)

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("should fail if body has unknown property", func(t *testing.T) {
		var error map[string]string
        var id = 2
        var pid = httprouter.Param{Key: "pid", Value: fmt.Sprint(id)}

		// tests '/api/posts/:id/comments' route
		request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/api/posts/%d/comments/", id), strings.NewReader(`{"title":"omg kys","user":"ILoveDogs96"}`))
		response := httptest.NewRecorder()
		request.Header.Add("Content-Type", "application/json")

		AddComment(response, request, []httprouter.Param{pid})

		err := json.NewDecoder(response.Body).Decode(&error)
		checkError(t, err)

		got := error
		want := map[string]string{"error": "couldn't add comment"}

		assertStatusCode(t, response, http.StatusBadRequest)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
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
