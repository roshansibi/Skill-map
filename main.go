package main


import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "skill-map-home-page",
		})
		r.GET("/login", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "user-login-required",
			})
	})
	r.Run() 
}

