package api

import (
	"fmt"
	"helen-portfolio/backend/models"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	// Create a unique name for this in-memory database instance
	// Each test will get its own isolated database
	dbID := fmt.Sprintf("file:memdb%d?mode=memory&cache=shared", time.Now().UnixNano())

	db, err := gorm.Open(sqlite.Open(dbID), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	// Keep connection alive for the duration of the test
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		sqlDB.Close()
	})

	err = db.AutoMigrate(&models.BlogPost{}, &models.BlogComment{})
	if err != nil {
		t.Fatal(err)
	}

	return db
}

func setupRouter(handler *Handler) *gin.Engine {
	router := gin.New()

	// contact
	router.POST("/api/contact", handler.SendEmail)
	// blog
	router.GET("/api/blog-previews", handler.GetBlogPreview)
	router.GET("/api/blog/:id", handler.GetBlogContent)
	router.POST("/api/blog-comment", handler.CreateBlogComment)
	// admin
	router.POST("/api/blog-new-entry", handler.CreateBlogPost)
	router.GET("/api/get-unapproved-comments", handler.GetUnapprovedComments)

	return router
}
