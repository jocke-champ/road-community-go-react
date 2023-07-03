package handlers 

import (
	"backebd/modals"
	"net/http"

	"github.com/gin-gonit/git
)

// CREATE methods

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Add user to database

	c.JSON(http.StatusOK, user)
}func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Add post to database

	c.JSON(http.StatusOK, post)
}

func CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Add comment to database

	c.JSON(http.StatusOK, comment)
}

func CreateLocation(c *gin.Context) {
	var location models.Location
	if err := c.ShouldBindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	// TODO: Add location to database

	c.JSON(http.StatusOK, location)
}

// EDIT methods

func EditPost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Update post in database

	c.JSON(http.StatusOK, post)
}

func EditComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Update comment in database

	c.JSON(http.StatusOK, comment)
}


// DELETE emthods

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	// TODO: Delete post from database using id

	c.JSON(http.StatusOK, gin.H{"status": "post deleted"})
}

func DeleteComment(c *gin.Context) {
	id := c.Param("id")

	// TODO: Delete comment from database using id

	c.JSON(http.StatusOK, gin.H{"status": "comment deleted"})
}
