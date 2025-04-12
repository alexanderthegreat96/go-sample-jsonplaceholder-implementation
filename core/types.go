package core

type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

type Post struct {
	Id       int    `json:"id"`
	UserId   int    `json:"userId"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Comments []Comment
}

type Worker struct {
	WorkerId int
	Posts    []Post
}
