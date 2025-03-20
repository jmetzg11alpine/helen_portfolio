package api

import (
	"helen-portfolio/backend/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
)

var secretKey = []byte("secret")

func (h *Handler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role := "user"
	if req.Username == "admin" && req.Password == "admin" {
		role = "admin"
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": req.Username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString, "role": role})

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

	if err := h.db.Create(&blogPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) GetUnapprovedComments(c *gin.Context) {
	var comments []models.UnapprovedComments

	result := h.db.Model(&models.BlogComment{}).
		Select("blog_comments.*, blog_posts.title as title").
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

func (h *Handler) ApproveComment(c *gin.Context) {
	var req models.CommentIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := h.db.Model(&models.BlogComment{}).Where("id= ?", req.ID).Update("approved", true)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to approve comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment approved successfully"})

}

func (h *Handler) DeleteComment(c *gin.Context) {
	var req models.CommentIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := h.db.Delete(&models.BlogComment{}, req.ID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to approve comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
