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

	switch r.Method {
	case http.MethodGet:
		if r.URL.Path == "" {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string][]Post{"posts": db})
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

            for i := range db {
                if db[i].Id == id {
                    w.Header().Add("Content-Type", "application/json")
                    w.WriteHeader(http.StatusOK)
                    json.NewEncoder(w).Encode(map[string][]Comment{"comments": db[i].Comments})
                    return
                }
            }
        }

		id, err := strconv.Atoi(r.URL.Path)

		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "post not found"})
			return
		}

		for i := range db {
			if db[i].Id == id {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(db[i])
				return
			}
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "post not found"})
	case http.MethodDelete:
		id, err := strconv.Atoi(r.URL.Path)

		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "post not found"})
			return
		}

		for i := range db {
			if db[i].Id != id {
				posts = append(posts, db[i])
			} else {
				post = db[i]
			}
		}

		if post.Id != 0 {
			db = posts
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

		if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
			http.Error(w, "'Content-Type' header is not 'application/json'", http.StatusUnsupportedMediaType)
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

		p.Id = db[len(db)-1].Id + 1
		p.Upvotes = 1
		p.Comments = []Comment{}
		db = append(db, p)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(p)
	default:
		w.Header().Set("Allow", "GET, POST, DELETE")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleComments(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		id, err := strconv.Atoi(r.URL.Path)

		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "comments for post not found"})
			return
		}

		for i := range db {
			if db[i].Id == id && len(db[i].Comments) > 1 {
                w.Header().Add("Content-Type", "application/json")
                w.WriteHeader(http.StatusOK)
                json.NewEncoder(w).Encode(map[string][]Comment{"comments": db[i].Comments})
                return
			}
		}
        
        w.Header().Add("Content-Type", "application/json")
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{"error": "post not found"})
	default:
		w.Header().Set("Allow", "GET")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}