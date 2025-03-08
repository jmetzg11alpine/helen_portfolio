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
