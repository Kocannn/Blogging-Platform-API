package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"Kocannn/Blogging-Platform-API.git/src/model"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func GetPosts(c *gin.Context, db *gorm.DB) {
	var posts []model.Posts
	if err := db.Find(&posts).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch posts"});
		return
	}

	c.JSON(http.StatusOK, posts)
}

func GetPostsById(c *gin.Context, db *gorm.DB){
	id := c.Param("id")
	var post model.Posts
	if err := db.First(&post,"id = ?" ,id).Error; err != nil{
    c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
    return
  }
	c.JSON(http.StatusOK, post)

}

func AddPost(c *gin.Context, db *gorm.DB){
	var newPost model.Posts
	if err := c.ShouldBindBodyWithJSON(&newPost); err != nil {
		c.AbortWithStatus(400);
		log.Fatal(err)
	}
	log.Print("Reciving data ", newPost)
	newPost.CreatedAt = time.Now()
	newPost.UpdatedAt = time.Now()
	
	if err := db.Create(&newPost).Error; err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, newPost)
	fmt.Println("Successfully added post")
} 

func DeletePost(c *gin.Context, db *gorm.DB){
	id := c.Param("id")
	var post model.Posts
	if err := db.First(&post,"id = ?" ,id).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	db.Delete(&post)
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}

func UpdatePost(c *gin.Context, db *gorm.DB){
	id := c.Param("id")
	var post model.Posts
	if err := db.First(&post,"id = ?" ,id).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	if err := c.ShouldBindBodyWithJSON(&post); err != nil {
		c.AbortWithStatus(400);
		log.Fatal(err)
	}
	db.Save(&post)
	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
}