package main

import (
	"crudapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the CRUD API",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/books", routes.GetBooks)
	router.GET("/book", routes.GetBook)
	router.POST("/book", routes.CreateBook)   // takes book struct as body of POST request, checks if bookid already exists, if not, creates new book,else returns error
	router.PATCH("/book", routes.UpdateBook)  // takes book struct as body of PATCH request, checks if bookid exists, if yes, updates book,else returns error
	router.DELETE("/book", routes.DeleteBook) // takes book struct as body of DELETE request, checks if bookid exists, if yes, deletes book,else returns error

	router.Run() // listen and serve on
}
