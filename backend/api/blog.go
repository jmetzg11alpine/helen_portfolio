package api

import (
	"fmt"
	"helen-portfolio/backend/models"
	"net/http"

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

	fmt.Printf("Attempting to fetch blog post with ID: %s\n", idStr)

	var blogPost models.BlogPost
	if err := h.db.Preload("Comments", "approved = ?", true).First(&blogPost, idStr).Error; err != nil {
		fmt.Printf("Error fetching blog post: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch blog post"})
		return
	}

	fmt.Printf("ID: %s", idStr)

	c.JSON(http.StatusOK, blogPost)
}
