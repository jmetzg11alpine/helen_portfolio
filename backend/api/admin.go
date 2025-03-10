package api

import (
	"helen-portfolio/backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateBlogPost(c *gin.Context) {
	var req models.CreateBlogPostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blogPost := models.BlogPost{
		Title:     req.Title,
		SubTitle:  req.SubTitle,
		Content:   req.Content,
		CreatedAt: time.Now().Unix(),
	}

	if err := h.db.Create(&blogPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) GetUnapprovedComments(c *gin.Context) {
	var comments []models.UnapprovedComments

	result := h.db.Model(&models.BlogComment{}).
		Select("blog_comments.*, blog_posts.title as title, blog_posts.sub_title as sub_title").
		Joins("JOIN blog_posts ON blog_comments.blog_id = blog_posts.id").
		Where("blog_comments.approved = ?", false).
		Order("blog_comments.created_at desc").
		Find(&comments)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch unapproved comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}
