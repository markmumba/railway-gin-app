package main

import (
 "os"

 "github.com/gin-gonic/gin"
)

func main() {

 // Intialise the Gin router
 router := gin.Default()

  // Define a GET route
 router.GET("/", func(c *gin.Context) {
  // Sends "Hello World" to the user
  c.JSON(200, gin.H{
   "message": "Hello, World!",
  })
 })

 // IMPORTANT: Set up a dynamic PORT
 port := os.Getenv("PORT")

 if port == "" {
  port = "8080" // Default port for development
 }

 // Start the router
 router.Run(":" + port)
}