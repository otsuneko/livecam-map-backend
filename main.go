package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})
	http.ListenAndServe(":8080", r) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
