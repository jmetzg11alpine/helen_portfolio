package api

import (
	"helen-portfolio/backend/models"
	"net/http"
	"net/smtp"
	"os"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SendEmail(c *gin.Context) {
	var req models.ContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid request format"})
		return
	}

	// Validate required fields
	if req.Email == "" || req.Name == "" || req.Message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Missing required fields"})
		return
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	from := os.Getenv("GMAIL_USER")
	password := os.Getenv("GMAIL_PASS")

	// Check if environment variables are set
	if from == "" || password == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Email configuration error"})
		return
	}

	// Prepare email content
	var subject, text string
	if req.Message == "newsletter" {
		text = "add " + req.Name + " to news letter list"
		subject = "News Letter - " + req.Name
	} else {
		text = req.Message
		subject = "Message from " + req.Name
	}

	// Compose email message with proper headers
	msg := []byte("From: " + from + "\r\n" + // Changed from req.Email to from
		"Reply-To: " + req.Email + "\r\n" +
		"To: " + from + "\r\n" +
		"Subject: " + subject + "\r\n\r\n" +
		text)

	// Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send email with detailed error logging
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{from}, msg) // Changed from req.Email to from
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to send email",
			"error":   err.Error(), // Include error details in development
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Email sent successfully!"})
}
