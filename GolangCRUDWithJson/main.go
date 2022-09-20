package main

import (
	"log"

	"example.com/demo/services"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.SetTrustedProxies([]string{"192.168.99.34"})
	r.GET("/GetAllBooks", func(c *gin.Context) {
		data, err := services.GetAllBooks()
		c.JSON(200, data)

		if err != nil {
			log.Fatal("Error when opening file: ", err)
		}
	})
	r.Run() // listen and serve on 0.0.0.0:8080
	// r.POST("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "nnnnn",
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080
}
