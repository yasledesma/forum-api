package main

import (
	// "encoding/json"
	// "fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	// "strings"
	"testing"
)

// func TestGetAllPosts(t *testing.T) {
// 	t.Run("should return all posts", func(t *testing.T) {
// 		var p []Post
// 		var posts = map[string][]Post{"posts": p}
//
// 		// tests '/api/posts' route
// 		request, _ := http.NewRequest(http.MethodGet, "/api/posts/", nil)
// 		response := httptest.NewRecorder()
//
// 		GetPosts(response, request, nil)
//
// 		err := json.NewDecoder(response.Body).Decode(&posts)
// 		checkError(t, err)
//
// 		got := posts["posts"]
// 		want := db.Posts
//
// 		assertStatusCode(t, response, http.StatusOK)
// 		assertSliceResponse(t, got, want)
// 	})
// }
//
// func TestCreatePost(t *testing.T) {
// 	t.Run("should succeed on post requests", func(t *testing.T) {
// 		var p Post
// 		body := `{"title":"Bobby","link":"https://beautyful-pictures.webs.com/photos/Cute-Puppies/puppy-lab-HD-wallpaper.jpg","username":"ILoveDogs96"}`
//
// 		// tests '/api/posts' route
// 		request, _ := http.NewRequest(http.MethodPost, "/api/posts/", strings.NewReader(body))
// 		response := httptest.NewRecorder()
// 		request.Header.Add("Content-Type", "application/json")
//
// 		AddPost(response, request, nil)
//
// 		err := json.NewDecoder(response.Body).Decode(&p)
// 		checkError(t, err)
//
// 		got := p
// 		want := db.Posts[len(db.Posts)-1]
//
// 		assertStatusCode(t, response, http.StatusCreated)
//
// 		if got != want {
// 			t.Errorf("got %v, want %v", got, want)
// 		}
// 	})
//
// 	t.Run("should fail if body has unknown property", func(t *testing.T) {
// 		var error map[string]string
// 		body := `{"title":"Bobby","link":"https://beautyful-pictures.webs.com/photos/Cute-Puppies/puppy-lab-HD-wallpaper.jpg","user":"ILoveDogs96"}`
//
// 		// tests '/api/posts' route
// 		request, _ := http.NewRequest(http.MethodPost, "/api/posts/", strings.NewReader(body))
// 		response := httptest.NewRecorder()
// 		request.Header.Add("Content-Type", "application/json")
//
// 		AddPost(response, request, nil)
//
// 		err := json.NewDecoder(response.Body).Decode(&error)
// 		checkError(t, err)
//
// 		got := err
// 		want := map[string]string{"error": "couldn't add post"}
//
// 		assertStatusCode(t, response, http.StatusBadRequest)
//
// 		if !reflect.DeepEqual(got, want) {
// 			t.Errorf("got %v, want %v", got, want)
// 		}
// 	})
// }
//
// func TestGetSinglePost(t *testing.T) {
// 	t.Run("should successfully return a single post", func(t *testing.T) {
// 		var p Post
// 		var id = 2
//
// 		// tests '/api/posts/:id' route
// 		request, _ := http.NewRequest(http.MethodGet, fmt.Sprint(id), nil)
// 		response := httptest.NewRecorder()
//
//         // same here, not passing params
// 		GetPost(response, request, make(httprouter.Params, id))
//
// 		err := json.NewDecoder(response.Body).Decode(&p)
// 		checkError(t, err)
//
// 		got := p
// 		want := db.Posts[id-1]
//
// 		assertStatusCode(t, response, http.StatusOK)
// 		if !reflect.DeepEqual(got, want) {
// 			t.Errorf("got %v, want %v", got, want)
// 		}
// 	})
//
// 	t.Run("should fail with non-existent post id", func(t *testing.T) {
// 		var error map[string]string
// 		var id = 483
//         var params = make(httprouter.Params, id)
//
// 		// tests '/api/posts/:id' route
// 		request, _ := http.NewRequest(http.MethodGet, fmt.Sprint(id), nil)
// 		response := httptest.NewRecorder()
//
//         // same here, not passing params
// 		GetPost(response, request, params)
//
// 		err := json.NewDecoder(response.Body).Decode(&error)
// 		checkError(t, err)
//
// 		got := error
// 		want := map[string]string{"error": "post not found"}
//
// 		assertStatusCode(t, response, http.StatusNotFound)
// 		if !reflect.DeepEqual(got, want) {
// 			t.Errorf("got %v, want %v", got, want)
// 		}
// 	})
// }
//
// func TestDeleteSinglePost(t *testing.T) {
// 	t.Run("should successfully delete a single post", func(t *testing.T) {
// 		var p Post
// 		var id = 1 
//         
// 		// tests '/api/posts/:id' route
// 		request, _ := http.NewRequest(http.MethodDelete, fmt.Sprint(id), nil)
// 		response := httptest.NewRecorder()
//
//         // maybe is not passing because of the params order, not being named, etc
//         // it fails because params are not being passed successfully
// 		DeletePost(response, request, make(httprouter.Params, id))
//
// 		err := json.NewDecoder(response.Body).Decode(&p)
// 		checkError(t, err)
//
// 		got := p.Id
// 		want := id
//
// 		assertStatusCode(t, response, http.StatusOK)
// 		if got != want {
// 			t.Errorf("got %v, want %v", got, want)
// 		}
// 	})
//
// 	t.Run("should fail with non-existent post id", func(t *testing.T) {
// 		var error map[string]string
// 		var id = 483
//
// 		// tests '/api/posts/:id' route
// 		request, _ := http.NewRequest(http.MethodDelete, fmt.Sprint(id), nil)
// 		response := httptest.NewRecorder()
//
//         // same here, not passing params
// 		DeletePost(response, request, make(httprouter.Params, id))
//
// 		err := json.NewDecoder(response.Body).Decode(&error)
// 		checkError(t, err)
//
// 		got := error
// 		want := map[string]string{"error": "can't delete post"}
//
// 		assertStatusCode(t, response, http.StatusNotFound)
// 		if !reflect.DeepEqual(got, want) {
// 			t.Errorf("got %v, want %v", got, want)
// 		}
// 	})
// }
//
// func TestGetAllCommentsForPost(t *testing.T) {
// 	t.Run("should return all comments for post", func(t *testing.T) {
// 		var c []Comment
// 		var comments = map[string][]Comment{"comments": c}
// 		var id = 1
//
// 		// tests '/api/posts/:id/comments' route
// 		request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%d/comments", id), nil)
// 		response := httptest.NewRecorder()
//
//         // same here, not passing params
// 		GetComments(response, request, make(httprouter.Params, id))
//
// 		err := json.NewDecoder(response.Body).Decode(&comments)
// 		checkError(t, err)
//
// 		got := comments["comments"][id - 1].PostId
// 		want := db.Comments[id - 1].PostId
//
// 		assertStatusCode(t, response, http.StatusOK)
// 		if got != want {
// 			t.Errorf("got %v, want %v", got, want)
// 		}
// 	})
//
//     
// 	t.Run("should fail with non-existent post id", func(t *testing.T) {
// 		var error map[string]string
// 		var id = 483
//
// 		// tests '/api/posts/:id/comments' route
// 		request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%d/comments", id), nil)
// 		response := httptest.NewRecorder()
//
//         // same here, not passing params
// 		GetComments(response, request, make(httprouter.Params, id))
//
// 		err := json.NewDecoder(response.Body).Decode(&error)
// 		checkError(t, err)
//
// 		got := error
// 		want := map[string]string{"error": "comments not found"}
//
// 		assertStatusCode(t, response, http.StatusNotFound)
// 		if !reflect.DeepEqual(got, want) {
// 			t.Errorf("got %v, want %v", got, want)
// 		}
// 	})
//
//     
// 	t.Run("should fail for posts with no comments", func(t *testing.T) {
// 		var error map[string]string
// 		var id = 2
//
// 		// tests '/api/posts/:id/comments' route
// 		request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%d/comments", id), nil)
// 		response := httptest.NewRecorder()
//
//         // same here, not passing params
// 		GetComments(response, request, make(httprouter.Params, id))
//
// 		err := json.NewDecoder(response.Body).Decode(&error)
// 		checkError(t, err)
//
// 		got := error
// 		want := map[string]string{"error": "comments not found"}
//
// 		assertStatusCode(t, response, http.StatusNotFound)
// 		if !reflect.DeepEqual(got, want) {
// 			t.Errorf("got %v, want %v", got, want)
// 		}
// 	})
// }
//
//
// func TestCreateComment(t *testing.T) {
// 	t.Run("should succeed on adding a comment", func(t *testing.T) {
// 		var c Comment 
//         body := `{"text":"so baby!","username":"ILoveDogs96"}`
//         id := 1
//
// 		// tests '/api/posts/:id/comments' route
// 		request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/api/posts/%d/comments", id), strings.NewReader(body))
// 		response := httptest.NewRecorder()
// 		request.Header.Add("Content-Type", "application/json")
//
//         // same here, not passing params
// 		AddComment(response, request, make(httprouter.Params, id))
//
// 		err := json.NewDecoder(response.Body).Decode(&c)
// 		checkError(t, err)
//
// 		got := c 
// 		want := db.Comments[c.Id - 1]
//
// 		assertStatusCode(t, response, http.StatusCreated)
//
// 		if got != want {
// 			t.Errorf("got %v, want %v", got, want)
// 		}
// 	})
//
// 	t.Run("should fail if body has unknown property", func(t *testing.T) {
// 		var error map[string]string
//         var id = 2
// 		body := `{"title":"omg kys","user":"ILoveDogs96"}`
//
// 		// tests '/api/posts/:id/comments' route
// 		request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%d/comments", id), strings.NewReader(body))
// 		response := httptest.NewRecorder()
// 		request.Header.Add("Content-Type", "application/json")
//
// 		AddPost(response, request, nil)
//         fmt.Println(response)
//
// 		err := json.NewDecoder(response.Body).Decode(&error)
// 		checkError(t, err)
//
// 		got := error
// 		want := map[string]string{"error": "couldn't add comment"}
//
// 		assertStatusCode(t, response, http.StatusBadRequest)
//
// 		if !reflect.DeepEqual(got, want) {
// 			t.Errorf("got %v, want %v", got, want)
// 		}
// 	})
// }
//
//
// func TestUpdateComment(t *testing.T) {
// 	t.Run("should update existing comment", func(t *testing.T) {
// 		var c Comment 
//         body := `{"text":"so baby! edit: change my mind, it's a cat"}`
//         postId := 1
//         commentId := 2
//
//         // tests '/api/posts/:postId/comments/:commentId' route
// 		request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%d/comments/%d", postId, commentId), strings.NewReader(body))
// 		response := httptest.NewRecorder()
// 		request.Header.Add("Content-Type", "application/json")
//
//         // same here, not passing params
// 		UpdateComment(response, request, make(httprouter.Params, postId, commentId))
//
// 		err := json.NewDecoder(response.Body).Decode(&c)
// 		checkError(t, err)
//
// 		got := c.Text
// 		want := db.Comments[c.Id - 1].Text
//
// 		assertStatusCode(t, response, http.StatusOK)
//
// 		if got != want {
// 			t.Errorf("got %v, want %v", got, want)
// 		}
// 	})
//
// 	t.Run("should fail if comment doesn't exist", func(t *testing.T) {
// 		var error map[string]string
//         body := `{"text":"so baby! edit: change my mind, it's a cat"}`
//         postId := 1
//         commentId := 234
//
//         // tests '/api/posts/:postId/comments/:commentId' route
// 		request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%d/comments/%d", postId, commentId), strings.NewReader(body))
// 		response := httptest.NewRecorder()
// 		request.Header.Add("Content-Type", "application/json")
//
//         // same here, not passing params
// 		UpdateComment(response, request, make(httprouter.Params, postId, commentId))
//
// 		err := json.NewDecoder(response.Body).Decode(&error)
// 		checkError(t, err)
//
// 		got := error
// 		want := map[string]string{"error": "couldn't edit comment"}
//
// 		assertStatusCode(t, response, http.StatusBadRequest)
//
// 		if !reflect.DeepEqual(got, want) {
// 			t.Errorf("got %v, want %v", got, want)
// 		}
// 	})
//
//     
// 	t.Run("should fail if post doesn't exist", func(t *testing.T) {
// 		var error map[string]string
//         body := `{"text":"so baby! edit: just joking!"}`
//         postId := 234
//         commentId := 2
//
//         // tests '/api/posts/:postId/comments/:commentId' route
// 		request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%d/comments/%d", postId, commentId), strings.NewReader(body))
// 		response := httptest.NewRecorder()
// 		request.Header.Add("Content-Type", "application/json")
//
//         // same here, not passing params
// 		UpdateComment(response, request, make(httprouter.Params, postId, commentId))
//         fmt.Println(response)
//
// 		err := json.NewDecoder(response.Body).Decode(&error)
// 		checkError(t, err)
//
// 		got := error
// 		want := map[string]string{"error": "post not found"}
//
// 		assertStatusCode(t, response, http.StatusBadRequest)
//
// 		if !reflect.DeepEqual(got, want) {
// 			t.Errorf("got %v, want %v", got, want)
// 		}
// 	})
// }

func assertSliceResponse(t *testing.T, got, want []Post) {
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
