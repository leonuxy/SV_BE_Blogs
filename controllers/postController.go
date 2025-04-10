package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"sv_be_posts/models"

	"github.com/gin-gonic/gin"
)

func validatePostInput(post *models.Post, c *gin.Context) bool {
	if len(post.Title) < 20 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title must be at least 20 characters"})
		return false
	}
	if len(post.Content) < 200 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Content must be at least 200 characters"})
		return false
	}
	if len(post.Category) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category must be at least 3 characters"})
		return false
	}
	status := strings.ToLower(post.Status)
	if status != "publish" && status != "draft" && status != "thrash" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status must be 'publish', 'draft', or 'thrash'"})
		return false
	}
	return true
}

func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !validatePostInput(&post, c) {
		return
	}
	if err := models.Create(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func GetAllPostsWithPagination(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Param("limit"))
	offset, _ := strconv.Atoi(c.Param("offset"))

	posts, err := models.GetPaginated(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func GetPostByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	post, err := models.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, post)
}

func UpdatePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !validatePostInput(&post, c) {
		return
	}
	post.ID = id
	if err := models.Update(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func DeletePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := models.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
