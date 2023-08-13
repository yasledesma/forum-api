package main

type Database struct {
	Posts    []Post    `json:"posts"`
	Comments []Comment `json:"comments"`
}

type Post struct {
	Id       int    `json:"id"`
	Upvotes  int    `json:"upvotes"`
	Title    string `json:"title"`
	Link     string `json:"link"`
	Username string `json:"username"`
}

type Comment struct {
	Id       int    `json:"id"`
	PostId   int    `json:"post_id"`
	Upvotes  int    `json:"upvotes"`
	Text     string `json:"text"`
	Username string `json:"username"`
}
