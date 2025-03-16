package api

import (
	"bytes"
	"encoding/json"
	"helen-portfolio/backend/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestCreateBlogPost(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db := setupTestDB(t)
	handler := NewHandler(db)
	router := setupRouter(handler)

	blogPost := models.CreateBlogPostRequest{
		Title:    "Test Blog Post",
		SubTitle: "A post for testing",
		Content:  "This is a detailed content of the test blog post.",
	}

	jsonData, err := json.Marshal(blogPost)
	if err != nil {
		t.Fatalf("Failed to marshal request: %v", err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/blog-new-entry", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}

	var savedPost models.BlogPost
	result := db.First(&savedPost, "title = ?", blogPost.Title)
	if result.Error != nil {
		t.Errorf("Failed to find saved post: %v", result.Error)
	}

	if savedPost.Title != blogPost.Title {
		t.Errorf("Expected title %s, got %s", blogPost.Title, savedPost.Title)
	}
	if savedPost.SubTitle != blogPost.SubTitle {
		t.Errorf("Expected subtitle %s, got %s", blogPost.SubTitle, savedPost.SubTitle)
	}
	if savedPost.Content != blogPost.Content {
		t.Errorf("Expected content %s, got %s", blogPost.Content, savedPost.Content)
	}
}

func TestCreateBlogPost_InvalidData(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db := setupTestDB(t)
	handler := NewHandler(db)
	router := setupRouter(handler)

	invalidBlogPost := struct {
		Title string `json:"title"`
	}{
		Title: "Just a title",
	}

	jsonData, err := json.Marshal(invalidBlogPost)
	if err != nil {
		t.Fatalf("Failed to marshal request: %v", err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/blog-new-entry", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}

	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if _, exists := response["error"]; !exists {
		t.Errorf("Expected error message in response")
	}
}

func TestGetUnapprovedComments(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db := setupTestDB(t)
	handler := NewHandler(db)
	router := setupRouter(handler)

	// Create test blog post
	blogPost := models.BlogPost{
		Title:     "Test Blog Post",
		SubTitle:  "A post for testing",
		Content:   "This is test content",
		CreatedAt: time.Now().Unix(),
	}
	if err := db.Create(&blogPost).Error; err != nil {
		t.Fatalf("Failed to create test blog post: %v", err)
	}

	// Create two unapproved comments
	unapprovedComment1 := models.BlogComment{
		BlogID:    blogPost.ID,
		Name:      "Test User 1",
		Content:   "First unapproved comment",
		Approved:  false,
		CreatedAt: time.Now().Unix(),
	}
	unapprovedComment2 := models.BlogComment{
		BlogID:    blogPost.ID,
		Name:      "Test User 2",
		Content:   "Second unapproved comment",
		Approved:  false,
		CreatedAt: time.Now().Unix() - 100,
	}

	approvedComment := models.BlogComment{
		BlogID:    blogPost.ID,
		Name:      "Approved User",
		Content:   "This comment is approved",
		Approved:  true,
		CreatedAt: time.Now().Unix(),
	}

	if err := db.Create(&unapprovedComment1).Error; err != nil {
		t.Fatalf("Failed to create unapproved comment 1: %v", err)
	}
	if err := db.Create(&unapprovedComment2).Error; err != nil {
		t.Fatalf("Failed to create unapproved comment 2: %v", err)
	}
	if err := db.Create(&approvedComment).Error; err != nil {
		t.Fatalf("Failed to create approved comment: %v", err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/get-unapproved-comments", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Parse response
	var comments []models.UnapprovedComments
	if err := json.Unmarshal(w.Body.Bytes(), &comments); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Should return exactly 2 unapproved comments
	if len(comments) != 2 {
		t.Errorf("Expected 2 comments, got %d", len(comments))
	}

	// Comments should be in descending order by creation time (newest first)
	if len(comments) >= 2 {
		if comments[0].CreatedAt <= comments[1].CreatedAt {
			t.Errorf("Comments not ordered by created_at desc, first: %d, second: %d",
				comments[0].CreatedAt, comments[1].CreatedAt)
		}
	}

	// Check that blog post details are included
	for _, comment := range comments {
		if comment.Title != blogPost.Title {
			t.Errorf("Expected blog title %s, got %s", blogPost.Title, comment.Title)
		}
		if comment.SubTitle != blogPost.SubTitle {
			t.Errorf("Expected blog subtitle %s, got %s", blogPost.SubTitle, comment.SubTitle)
		}
	}
}

func TestGetUnapprovedComments_NoComments(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db := setupTestDB(t)
	handler := NewHandler(db)
	router := setupRouter(handler)

	// Create a blog post but no comments
	blogPost := models.BlogPost{
		Title:     "Test Blog Post",
		SubTitle:  "A post for testing",
		Content:   "This is test content",
		CreatedAt: time.Now().Unix(),
	}
	if err := db.Create(&blogPost).Error; err != nil {
		t.Fatalf("Failed to create test blog post: %v", err)
	}

	// Make request to get unapproved comments
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/get-unapproved-comments", nil)
	router.ServeHTTP(w, req)

	// Check response status
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Parse response
	var comments []models.UnapprovedComments
	if err := json.Unmarshal(w.Body.Bytes(), &comments); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Should return an empty array
	if len(comments) != 0 {
		t.Errorf("Expected 0 comments, got %d", len(comments))
	}
}
