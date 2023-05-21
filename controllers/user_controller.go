package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ogs-create-api/models"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
)

func GetActiveUsers(c echo.Context) error {
	var users []models.ActiveUser

	client := resty.New()

	resp, err := client.R().Get("https://gorest.co.in/public/v2/users")
	if err != nil {
		return err
	}

	json.Unmarshal(resp.Body(), &users)

	var activeUsers []models.ActiveUser

	for _, user := range users {

		if user.Status == "active" {
			activeUsers = append(activeUsers, user)
		}
	}

	return c.JSON(http.StatusOK, activeUsers)
}

func GetUserPost(c echo.Context) error {
	var users []models.UserPost
	var posts []models.Post
	var comments []models.Comment

	client := resty.New()

	resp_users, err := client.R().Get("https://gorest.co.in/public/v2/users")
	if err != nil {
		return err
	}
	json.Unmarshal(resp_users.Body(), &users)

	resp_posts, err := client.R().Get("https://gorest.co.in/public/v2/posts")
	if err != nil {
		return err
	}
	json.Unmarshal(resp_posts.Body(), &posts)

	resp_comments, err := client.R().Get("https://gorest.co.in/public/v2/comments")
	if err != nil {
		return err
	}
	json.Unmarshal(resp_comments.Body(), &comments)

	var userPosts []models.UserPost

	for _, user := range users {
		if user.Status == "active" {
			userPost := models.UserPost{
				Id:          user.Id,
				Name:        user.Name,
				Email:       user.Email,
				Gender:      user.Gender,
				Status:      user.Status,
				Post_Amount: user.Post_Amount,
				Posts:       nil,
			}

			for _, post := range posts {
				if post.User_id == user.Id {
					post.Comments = nil
					for _, comment := range comments {
						if comment.Post_id == post.Id {
							post.Comments = append(post.Comments, comment)
							post.Comment_Amount++
						}
					}
					userPost.Posts = append(userPost.Posts, post)
				}
			}

			userPost.Post_Amount = len(userPost.Posts)
			userPosts = append(userPosts, userPost)
		}
	}

	fmt.Println(users)
	fmt.Println(posts)

	return c.JSON(http.StatusOK, userPosts)
}
