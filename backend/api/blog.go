package api

import (
	"fmt"
	"helen-portfolio/backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetBlogPreview(c *gin.Context) {
	var blogPreviews []models.BlogPostPreview

	result := h.db.Model(&models.BlogPost{}).
		Select("blog_posts.id, blog_posts.title, blog_posts.sub_title, blog_posts.created_at, COUNT(blog_comments.id) as comment_count").
		Joins("LEFT JOIN blog_comments ON blog_comments.blog_id = blog_posts.id").
		Group("blog_posts.id").
		Find(&blogPreviews)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch blog previews"})
		return
	}

	c.JSON(http.StatusOK, blogPreviews)
}

func (h *Handler) GetBlogContent(c *gin.Context) {
	idStr := c.Param("id")

	var blogPost models.BlogPost
	if err := h.db.Preload("Comments", "approved = ?", true).First(&blogPost, idStr).Error; err != nil {
		fmt.Printf("Error fetching blog post: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch blog post"})
		return
	}

	c.JSON(http.StatusOK, blogPost)
}

func (h *Handler) CreateBlogComment(c *gin.Context) {
	var req models.CreateBlogCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	comment := models.BlogComment{
		Name:      req.Name,
		Content:   req.Comment,
		BlogID:    req.BlogID,
		CreatedAt: time.Now().Unix(),
		Approved:  false,
	}

	if err := h.db.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.Status(http.StatusCreated)
}
