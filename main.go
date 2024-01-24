package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	// TODO: make an api middleware
	router.GET("/api/posts/", GetPosts)
	router.POST("/api/posts/", AddPost)
	router.GET("/api/posts/:pid/", GetPost)
	router.DELETE("/api/posts/:pid/", DeletePost)
	router.POST("/api/posts/:pid/comments/", AddComment)
	router.GET("/api/posts/:pid/comments/", GetComments)
	router.POST("/api/posts/:pid/comments/:cid", UpdateComment)

	fmt.Println("Server running on http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

