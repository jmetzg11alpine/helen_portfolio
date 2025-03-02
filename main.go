package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"helen-portfolio/backend/database"
	"helen-portfolio/backend/routes"
)

func main() {

	database.Connect()
	fmt.Println("Database connected successfully!")
	r := gin.Default()

	routes.SetupAPIRoutes(r)
	routes.SetupStaticRoutes(r)

	// Start the server
	r.Run(":3000")
}
