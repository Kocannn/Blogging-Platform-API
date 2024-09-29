package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type blog struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Tags      []string `json:"tags"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

var posts = []blog{
	{Id: "1", Title: "Seni Bersikap Bodo Amat", Content: "Seni Bersikap Bodo Amat", Tags: []string{"Seni", "Bersikap", "Bodo", "Amat"}, CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

func main() {
	PORT := (":8080")
	r := gin.Default()
	
	r.GET("/posts", getAllPosts)
	r.GET("/posts/:id", getPost)
	r.POST("posts", addNewPost)
	r.PUT("posts/:id", updatePost)
	r.DELETE("posts/:id", deletePost)

	r.Run(PORT)
}

func getPost(c *gin.Context){
	id := c.Param("id")
	for _, post := range posts{
		if post.Id == id{
			c.JSON(200, post)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "post not found"})
}

func addNewPost(c * gin.Context){
	var newPosts *blog
	if err := c.BindJSON(&newPosts); err != nil{
		c.AbortWithStatus(400)
		log.Fatal(err)
	}
	newPosts.CreatedAt = time.Now()
	newPosts.UpdatedAt = time.Now()

	posts = append(posts, *newPosts)
	c.JSON(200, newPosts)
}

func getAllPosts(c *gin.Context){
	c.JSON(200, posts)
}

func updatePost(c *gin.Context){
	id := c.Param("id")

	var updatePost *blog
	if err := c.BindJSON(&updatePost); err != nil{
		c.AbortWithStatus(400)
		log.Fatal(err)
	}

	for i, post := range posts{
		if post.Id == id{
			createdAtOriginal := posts[i].CreatedAt
			posts[i] = *updatePost
			posts[i].UpdatedAt = time.Now()
			posts[i].CreatedAt = createdAtOriginal
			c.JSON(200, updatePost)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "post not found"})

}

func deletePost(c *gin.Context){
	id := c.Param("id")

	var deletePost *blog
	if err := c.BindJSON(&deletePost); err != nil{
		c.AbortWithStatus(400)
		log.Fatal(err)
	}

	for i, post := range posts{
		if post.Id == id{
			posts = append(posts[:i], posts[i+1:]...)
			c.JSON(200, gin.H{"message": "post deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "post not found"})

}