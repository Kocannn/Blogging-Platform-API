package main

import (
	"Kocannn/Blogging-Platform-API.git/src/config"
	"Kocannn/Blogging-Platform-API.git/src/controllers"

	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()

	db := config.ConnectDB()

	config.MigrateDB(db)

	r.GET("/posts", func(c *gin.Context) {
		controllers.GetPosts(c, db)
	})	
	
	r.GET("/posts/:id", func(c *gin.Context) {
		controllers.GetPostsById(c, db)
	})

	r.POST("/posts", func(c *gin.Context){
		controllers.AddPost(c, db)
	})
	
	r.PUT("/posts/:id", func(c *gin.Context){
		controllers.UpdatePost(c, db)
	})

	r.DELETE("/posts/:id", func(c *gin.Context){
		controllers.DeletePost(c, db)
	})
	
	r.Run(":8080")
	// r.Run(PORT)
}

