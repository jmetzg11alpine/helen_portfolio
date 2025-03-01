package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Enable CORS for local development
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Allow local frontend in dev
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// API routes under `/api`
	api := r.Group("/api")
	{
		api.GET("/hello", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Hello from Go!"})
		})
	}

	// Serve static files, but avoid conflicting with `/api`
	r.Use(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.Next() // Let API routes handle it
			return
		}

		// Try to serve static files
		filePath := "./frontend/build" + c.Request.URL.Path
		if _, err := http.Dir("./frontend/build").Open(c.Request.URL.Path); err == nil {
			c.File(filePath)
			c.Abort()
			return
		}

		// If not a file, serve index.html (for SvelteKit routing)
		c.File("./frontend/build/index.html")
		c.Abort()
	})

	// Start the server
	r.Run(":3000")
}
