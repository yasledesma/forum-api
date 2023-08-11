package main

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
