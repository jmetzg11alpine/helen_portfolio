// func AdminMiddleware() gin.HandlerFunc {
//     return func(c *gin.Context) {
//         tokenString := c.GetHeader("Authorization")
//         if tokenString == "" {
//             c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
//             c.Abort()
//             return
//         }

//         token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//             return secretKey, nil
//         })

//         if err != nil || !token.Valid {
//             c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
//             c.Abort()
//             return
//         }

//         claims, ok := token.Claims.(jwt.MapClaims)
//         if !ok || claims["role"] != "admin" {
//             c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
//             c.Abort()
//             return
//         }

//         c.Next()
//     }
// }

// protecting routes

// r := gin.Default()
// r.POST("/login", LoginHandler)

// admin := r.Group("/admin")
// admin.Use(AdminMiddleware())
// admin.GET("/dashboard", func(c *gin.Context) {
//     c.JSON(http.StatusOK, gin.H{"message": "Welcome Admin!"})
// })

// r.Run(":8080")
