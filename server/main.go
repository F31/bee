package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, this is a Gin microservice!",
		})
	})

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
