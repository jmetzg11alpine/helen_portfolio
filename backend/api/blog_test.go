package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"helen-portfolio/backend/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetBlogPreview(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db := setupTestDB(t)

	handler := NewHandler(db)
	router := setupRouter(handler)

	testPost1 := models.BlogPost{
		Title:    "Test Post 1",
		SubTitle: "First test post",
		Content:  "This is the content of the first test post",
	}

	testPost2 := models.BlogPost{
		Title:    "Test Post 2",
		SubTitle: "Second test post",
		Content:  "This is the content of the second test post",
	}

	db.Create(&testPost1)
	db.Create(&testPost2)

	comment1 := models.BlogComment{
		BlogID:   testPost1.ID,
		Name:     "Commenter 1",
		Content:  "Great post!",
		Approved: true,
	}

	comment2 := models.BlogComment{
		BlogID:   testPost1.ID,
		Name:     "Commenter 2",
		Content:  "I learned a lot.",
		Approved: true,
	}

	db.Create(&comment1)
	db.Create(&comment2)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/blog-previews", nil)

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var response []models.BlogPostPreview
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	if len(response) != 2 {
		t.Errorf("Expected 2 blog posts, got %d", len(response))
	}

	var firstPost models.BlogPostPreview
	for _, post := range response {
		if post.ID == testPost1.ID {
			firstPost = post
			break
		}
	}

	if firstPost.CommentCount != 2 {
		t.Errorf("Expected 2 comments for first post, got %d", firstPost.CommentCount)
	}
}

func TestGetBlogContent(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db := setupTestDB(t)
	handler := NewHandler(db)
	router := setupRouter(handler)

	testPost := models.BlogPost{
		Title:    "Test Blog Post",
		SubTitle: "A post for testing",
		Content:  "This is a detailed content of the test blog post.",
	}

	db.Create(&testPost)

	approvedComment := models.BlogComment{
		BlogID:   testPost.ID,
		Name:     "Approved User",
		Content:  "This is an approved comment.",
		Approved: true,
	}

	unapprovedComment := models.BlogComment{
		BlogID:   testPost.ID,
		Name:     "Unapproved User",
		Content:  "This comment should not appear.",
		Approved: false,
	}

	db.Create(&approvedComment)
	db.Create(&unapprovedComment)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/blog/%d", testPost.ID), nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var response models.BlogPost
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	if response.ID != testPost.ID {
		t.Errorf("Expected blog ID %d, got %d", testPost.ID, response.ID)
	}

	if response.Title != testPost.Title {
		t.Errorf("Expected title '%s', got '%s'", testPost.Title, response.Title)
	}

	if len(response.Comments) != 1 {
		t.Errorf("Expected 1 approved comment, got %d", len(response.Comments))
	}

	if len(response.Comments) > 0 && !response.Comments[0].Approved {
		t.Error("Unapproved comment was included in response")
	}

}

func TestCreateBlogComment(t *testing.T) {

	gin.SetMode(gin.TestMode)

	db := setupTestDB(t)
	handler := NewHandler(db)
	router := setupRouter(handler)

	testPost := models.BlogPost{
		Title:    "Test Post for Comments",
		SubTitle: "Testing comment creation",
		Content:  "This is a post that will receive comments",
	}
	db.Create(&testPost)

	commentRequest := models.CreateBlogCommentRequest{
		BlogID:  testPost.ID,
		Name:    "Test Commenter",
		Comment: "This is a test comment.",
	}

	jsonData, err := json.Marshal(commentRequest)
	if err != nil {
		t.Fatalf("Failed to marshal request: %v", err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/blog-comment", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}

	var comments []models.BlogComment
	db.Where("blog_id = ?", testPost.ID).Find(&comments)

	if len(comments) != 1 {
		t.Errorf("Expected 1 comment in database, got %d", len(comments))
	}

	if len(comments) > 0 {
		comment := comments[0]
		if comment.Name != commentRequest.Name {
			t.Errorf("Expected commenter name '%s', got '%s'", commentRequest.Name, comment.Name)
		}

		if comment.Content != commentRequest.Comment {
			t.Errorf("Expected comment content '%s', got '%s'", commentRequest.Comment, comment.Content)
		}

		if comment.Approved {
			t.Error("New comment should be unapproved by default")
		}
	}

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/blog-comment", bytes.NewBuffer([]byte(`{"invalid json"}`)))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d for invalid JSON, got %d", http.StatusBadRequest, w.Code)
	}
}
