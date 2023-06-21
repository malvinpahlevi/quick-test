package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("template/**")

	router.GET("/", func(c *gin.Context) {

		c.HTML(200, "index.html", gin.H{
			"msg": "success",
		})
	})

	router.Run(":8080")
}
