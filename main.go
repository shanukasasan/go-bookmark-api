package main

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Bookmark struct {
	ID    int    `json:"id"`
	Title string `json:"title" binding:"required"`
	URL   string `json:"url" binding:"required,url"`
}

var bookmarks = []Bookmark{
	{ID: 1, Title: "Google", URL: "https://www.google.com"},
	{ID: 2, Title: "GitHub", URL: "https://github.com"},
	{ID: 3, Title: "Stack Overflow", URL: "https://stackoverflow.com"},
}
var nextID = 4

func main() {

	router := gin.Default()

	// Read all bookmarks
	router.GET("/bookmarks", func(c *gin.Context) {
		c.JSON(200, bookmarks)
	})

	// Read one bookmark
	router.GET("/bookmarks/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		for _, bookmark := range bookmarks {
			if bookmark.ID == id {
				c.JSON(200, bookmark)
				return
			}
		}
		c.JSON(404, gin.H{"error": "Bookmark not found"})
	})

	// Update a bookmark
	router.PUT("/bookmarks/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var updatedBookmark Bookmark

		if err := c.ShouldBindJSON(&updatedBookmark); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		for i, bookmark := range bookmarks {
			if bookmark.ID == id {
				bookmarks[i] = updatedBookmark
				c.JSON(200, updatedBookmark)
				return
			}
		}
		c.JSON(404, gin.H{"error": "Bookmark not found"})
	})

	// Delete a bookmark

	router.DELETE("/bookmarks/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		for i, bookmark := range bookmarks {
			if bookmark.ID == id {
				bookmarks = append(bookmarks[:i], bookmarks[i+1:]...)
				c.JSON(200, gin.H{"message": "Bookmark deleted"})
				return
			}
		}
		c.JSON(404, gin.H{"error": "Bookmark not found"})
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.GET("hello/:name", func(c *gin.Context) {
		//Path parameter
		name := c.Param("name")

		//Query parameter
		loud := c.DefaultQuery("loud", "false")

		greeting := "Hello " + name

		if loud == "true" {
			greeting = strings.ToUpper(greeting)
		}
		c.JSON(200, gin.H{
			"message": greeting,
		})
	})

	router.POST("/bookmarks", func(c *gin.Context) {
		var newBookmark Bookmark

		//JSON BINDING
		if err := c.ShouldBindJSON(&newBookmark); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		newBookmark.ID = nextID
		nextID++
		bookmarks = append(bookmarks, newBookmark)
		c.JSON(201, newBookmark)
	})

	err := router.Run(":8080")
	if err != nil {
		println("Error starting server:", err)
	}
}
