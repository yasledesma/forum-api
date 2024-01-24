package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func deletePostComments(pid int) []Comment {
    var comments []Comment
    
    for i := range db.Comments {
        if db.Comments[i].PostId == pid {
            comments = append(comments, db.Comments[i])
        } 
    }

    return comments
}

func GetPosts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string][]Post{"posts": db.Posts})
}

func AddPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	// TODO: this is business logic. move it to a service.
	// check for empty field

	p.Id = db.Posts[len(db.Posts)-1].Id + 1 // save as next index
	p.Upvotes = 1
	db.Posts = append(db.Posts, p)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func GetPost(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("pid"))

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
}

func DeletePost(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var post Post
	var posts []Post

	id, err := strconv.Atoi(params.ByName("pid"))

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
            db.Comments = deletePostComments(i)
		}
	}

	if post.Id != 0 {
		db.Posts = posts
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(post)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "can't delete post"})
}

func GetComments(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var comments []Comment

	id, err := strconv.Atoi(params.ByName("pid"))

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
}

func AddComment(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var c Comment

	id, err := strconv.Atoi(params.ByName("pid"))

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

	var index int = 1

	if len(db.Comments) == 0 {
		index = 1
	} else {
        index = db.Comments[len(db.Comments) - 1].Id + 1 
    }

	c.Id = index
	c.PostId = id
	c.Upvotes = 1
	db.Comments = append(db.Comments, c)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

func UpdateComment(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var c Comment

	pid, err := strconv.Atoi(params.ByName("pid"))
    
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "post not found"})
		return
	}

	cid, err := strconv.Atoi(params.ByName("cid"))

	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "comment not found"})
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	for i := range db.Comments {
		if db.Comments[i].PostId == pid && db.Comments[i].Id == cid {
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

	db.Comments[cid] = c
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(c)
}
