package models

type ActiveUser struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
	Status string `json:"status"`
}

type Post struct {
	Id             int       `json:"id"`
	User_id        int       `json:"user_id"`
	Title          string    `json:"title"`
	Body           string    `json:"body"`
	Comment_Amount int       `json:"comment_amount"`
	Comments       []Comment `json:"comments"`
}

type Comment struct {
	Id      int    `json:"id"`
	Post_id int    `json:"post_id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Body    string `json:"body"`
}

type UserPost struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Gender      string `json:"gender"`
	Status      string `json:"status"`
	Post_Amount int    `json:"post_amount"`
	Posts       []Post `json:"posts"`
}
