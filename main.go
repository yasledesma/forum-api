package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.Handle(
        "/api/posts/",
		http.StripPrefix("/api/posts/", http.HandlerFunc(handlePosts)),
	)

	fmt.Println("Server running on http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

