package api

import (
	"helen-portfolio/backend/database"
	"helen-portfolio/backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) GetBlogPreview(c *gin.Context) {
	var blogPreviews []models.BlogPostPreview

	result := database.DB.Model(&models.BlogPost{}).
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

	if err := database.DB.Create(&blogPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}
