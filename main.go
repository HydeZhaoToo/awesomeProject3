package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	r.GET("/ping/:name", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": c.Param("name"),
		})
	})
	r.GET("/dsa",test)
	err := r.Run()
	if err != nil {
		return 
	}
}
func test(c *gin.Context)  {
	c.String(200,"ds'a'ds\n")
}