package database

import (
	"helen-portfolio/backend/models"
	"log"
	"time"

	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
	var count int64
	db.Model(&models.BlogPost{}).Count(&count)
	if count > 0 {
		log.Println("Database already seeded")
		return
	}

	log.Println("Seeding development database...")

	blogPost := models.BlogPost{
		Title:     "My First Blog Post",
		SubTitle:  "This is the sub title of my first blog post",
		Content:   "This is the content of my first blog post. It is longer than the sub title. Only for testing purposes.",
		CreatedAt: time.Now().Unix(),
	}
	db.Create(&blogPost)

	blogComment := models.BlogComment{
		Content:   "This is a comment on the blog post.",
		Name:      "John Doe",
		BlogID:    blogPost.ID,
		CreatedAt: time.Now().Unix(),
		Approved:  true,
	}
	db.Create(&blogComment)

	blogComment2 := models.BlogComment{
		Content:   "This is a comment is waiting for approval on the blog post.",
		Name:      "Mary Jane",
		BlogID:    blogPost.ID,
		CreatedAt: time.Now().Unix(),
		Approved:  false,
	}
	db.Create(&blogComment2)

	merch := models.Merch{
		Title:     "My First Merch",
		Price:     10.0,
		Quantity:  100,
		CreatedAt: time.Now().Unix(),
	}
	db.Create(&merch)

	merchHistory := models.MerchHistory{
		ItemName:  merch.Title,
		Price:     int(merch.Price),
		Quantity:  merch.Quantity,
		CreatedAt: time.Now().Unix(),
	}
	db.Create(&merchHistory)

	log.Println("Database seeded successfully!")
}
