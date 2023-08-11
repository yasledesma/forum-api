package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Post struct {
	Id       int       `json:"id"`
	Upvotes  int       `json:"upvotes"`
	Title    string    `json:"title"`
	Link     string    `json:"link"`
	Username string    `json:"username"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	Id       int    `json:"id"`
	Upvotes  int    `json:"upvotes"`
	Text     string `json:"text"`
	Username string `json:"username"`
}

var db = []Post{
	{
		Id:       1,
		Upvotes:  1,
		Title:    "My cat is the cutest!",
		Link:     "https://i.imgur.com/jseZqNK.jpg",
		Username: "alicia98",
		Comments: []Comment{
			{
				Id:       1,
				Upvotes:  2,
				Text:     "She's such a cutie! :3",
				Username: "raahi014",
			},
		},
	},
	{
		Id:       2,
		Upvotes:  -432,
		Title:    "Thomas Jefferson circa 2015",
		Link:     "https://i.redd.it/xn9auq3xdoa51.jpg",
		Username: "ZzturtleszZ",
		Comments: []Comment{},
	},
}

func main() {

	router := http.NewServeMux()

	router.HandleFunc("/api/posts", handlePosts)

	fmt.Println("Server running on http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func handlePosts(w http.ResponseWriter, r *http.Request) {
	if http.MethodGet == r.Method {
        fmt.Println(http.MethodGet, r.Method)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string][]Post{"posts": db})
	} else if http.MethodPost == r.Method {
        var p Post
        err := json.NewDecoder(r.Body).Decode(&p)

        if err != nil {
            w.Header().Add("Content-Type", "application/json")
            w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(map[string]string{"error": "Couldn't add post"})
            return
        }

        p.Id = len(db) + 1
        p.Upvotes = 1
        p.Comments = []Comment{}
        db = append(db, p)
        
        w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(p)

    } else {
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    }
}
