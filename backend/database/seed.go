package database

import (
	"helen-portfolio/backend/models"
	"log"
	"time"

	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
	var count int64
	db.Model(&models.User{}).Count(&count)
	if count > 0 {
		log.Println("Database already seeded")
		return
	}

	log.Println("Seeding development database...")

	adminUser := models.User{Email: "admin@example.com", Username: "admin", Role: "admin"}
	regularUser := models.User{Email: "user@example.com", Username: "user", Role: "user"}
	db.Create(&adminUser)
	db.Create(&regularUser)

	blogPost := models.BlogPost{
		Title:     "My First Blog Post",
		SubTitle:  "This is the sub title of my first blog post",
		Content:   "This is the content of my first blog post. It is longer than the sub title. Only for testing purposes.",
		CreatedAt: time.Now().Unix(),
	}
	db.Create(&blogPost)

	blogComment := models.BlogComment{
		Content:   "This is a comment on the blog post.",
		UserID:    regularUser.ID,
		BlogID:    blogPost.ID,
		CreatedAt: time.Now().Unix(),
	}
	db.Create(&blogComment)

	merch := models.Merch{
		Title:     "My First Merch",
		Price:     10.0,
		Quantity:  100,
		CreatedAt: time.Now().Unix(),
	}
	db.Create(&merch)

	merchHistory := models.MerchHistory{
		UserID:    regularUser.ID,
		ItemName:  merch.Title,
		Price:     int(merch.Price),
		Quantity:  merch.Quantity,
		CreatedAt: time.Now().Unix(),
	}
	db.Create(&merchHistory)

	log.Println("Database seeded successfully!")
}
