package controllers

import (
	"latihan/helper"
	"latihan/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InPost struct {
	Post *helper.Post
}

func (in *InPost) GetAllPosts(c *gin.Context) {
	if !middleware.Auth(c) {
		return
	}

	posts, err := in.Post.GetAllPosts("/posts")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func (in *InPost) GetPostById(c *gin.Context) {
	if !middleware.Auth(c) {
		return
	}
	id := c.Param("id")
	post, err := in.Post.GetPostById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}

func (in *InPost) CreateNewPost(c *gin.Context) {
	if !middleware.Auth(c) {
		return
	}
	var post helper.Post
	c.BindJSON(&post)

	err := in.Post.CreateNewPost(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}
