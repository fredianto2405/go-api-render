package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		currentTime := time.Now().Format(time.RFC3339)
		c.JSON(http.StatusOK, gin.H{
			"pong": currentTime,
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
