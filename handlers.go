package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func handlePosts(w http.ResponseWriter, r *http.Request) {
	var post Post
	var posts []Post
	var comments []Comment

	switch r.Method {
	case http.MethodGet:
		if r.URL.Path == "" {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string][]Post{"posts": db.Posts})
			return
		}

		if strings.Contains(r.URL.Path, "comments") {
			id, err := strconv.Atoi(r.URL.Path[:strings.Index(r.URL.Path, "/")])

			if err != nil {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]string{"error": "comments not found"})
				return
			}

			for i := range db.Comments {
				if db.Comments[i].PostId == id {
					comments = append(comments, db.Comments[i])
				}
			}

			if comments != nil {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(map[string][]Comment{"comments": comments})
				return
			}

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "comments not found"})
			return
		}

		id, err := strconv.Atoi(r.URL.Path)

		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "post not found"})
			return
		}

		for i := range db.Posts {
			if db.Posts[i].Id == id {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(db.Posts[i])
				return
			}
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "post not found"})
	case http.MethodDelete:
		var comments []Comment

		id, err := strconv.Atoi(r.URL.Path)

		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "post not found"})
			return
		}

		for i := range db.Posts {
			if db.Posts[i].Id != id {
				posts = append(posts, db.Posts[i])
			} else {
				post = db.Posts[i]
			}
		}

		for i := range db.Comments {
			if db.Comments[i].PostId != id {
				comments = append(comments, db.Comments[i])
			}
		}

		if post.Id != 0 {
			db.Posts = posts
			db.Comments = comments
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(post)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "can't delete post"})
	case http.MethodPost:
		var p Post
		var c Comment

		if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
			http.Error(w, "'Content-Type' header is not 'application/json'", http.StatusUnsupportedMediaType)
			return
		}

		if strings.Contains(r.URL.Path, "comments/") {
			id, err := strconv.Atoi(r.URL.Path[strings.LastIndexAny(r.URL.Path, "/")+1:])

			if err != nil {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]string{"error": "comment not found"})
				return
			}

			r.Body = http.MaxBytesReader(w, r.Body, 1048576)

			for i := range db.Comments {
				if db.Comments[i].Id == id {
					c = db.Comments[i]
				}
			}

			dec := json.NewDecoder(r.Body)
			dec.DisallowUnknownFields()
			err = dec.Decode(&c)

			if err != nil || c.Id < 1 {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"error": "couldn't edit comment"})
				return
			}

			db.Comments[id-1] = c
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(c)
			return
		}

		if strings.Contains(r.URL.Path, "/comments") {
			id, err := strconv.Atoi(r.URL.Path[:strings.Index(r.URL.Path, "/")])

			if err != nil {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]string{"error": "post not found"})
				return
			}

			r.Body = http.MaxBytesReader(w, r.Body, 1048576)

			dec := json.NewDecoder(r.Body)
			dec.DisallowUnknownFields()
			err = dec.Decode(&c)

			if err != nil {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"error": "couldn't add post"})
				return
			}

			var index int = 1

			if len(db.Comments) != 0 {
				index = len(db.Comments) - 1
			}

			c.Id = db.Comments[index].Id + 1
			c.PostId = id
			c.Upvotes = 1
			db.Comments = append(db.Comments, c)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(c)
			return
		}

		r.Body = http.MaxBytesReader(w, r.Body, 1048576)

		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		err := dec.Decode(&p)

		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "couldn't add post"})
			return
		}

		p.Id = db.Posts[len(db.Posts)-1].Id + 1
		p.Upvotes = 1
		db.Posts = append(db.Posts, p)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(p)
	default:
		w.Header().Set("Allow", "GET, POST, DELETE")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
