package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"crudapi/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

var database = ConnectToServer()

func GetBooks(c *gin.Context) {
	var books []models.Book
	collection := database.Collection("books")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var book models.Book
		cursor.Decode(&book)
		books = append(books, book)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": books})
}

func GetBook(c *gin.Context) {
	fmt.Println(c.Query("bookid"))
	collection := database.Collection("books")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var book models.Book
	// string to int
	bookid, err := strconv.Atoi(c.Query("bookid"))
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "bookid must be an integer"})
		return
	}

	filter := bson.M{"bookid": bookid}
	cursor, err := collection.Find(ctx, filter)
	// if more than one book with same bookid
	if cursor.RemainingBatchLength() > 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "more than one book with same bookid"})
		return
	}
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		cursor.Decode(&book)
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func CreateBook(c *gin.Context) {
	// POST request body
	var input models.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	collection := database.Collection("books")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// insert book
	// check if bookid already exists
	filter := bson.M{"bookid": input.BOOKID}
	cursor, err := collection.Find(ctx, filter)
	if cursor.RemainingBatchLength() > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bookid already exists"})
		return
	}
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	newBook := bson.M{"bookid": input.BOOKID, "title": input.Title, "author": input.Author, "year": input.Year}
	_, err = collection.InsertOne(ctx, newBook)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": newBook})

}

func UpdateBook(c *gin.Context) {
	// PUT request body
	var input models.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	collection := database.Collection("books")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// update book
	// check if bookid exists
	filter := bson.M{"bookid": input.BOOKID}
	cursor, err := collection.Find(ctx, filter)
	if cursor.RemainingBatchLength() == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bookid does not exist"})
		return
	}
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	// update book
	update := bson.M{"$set": bson.M{"title": input.Title, "author": input.Author, "year": input.Year}}

	// update book
	filter = bson.M{"bookid": input.BOOKID}

	collection.FindOneAndUpdate(ctx, filter, update)
	c.JSON(http.StatusOK, gin.H{"data": input})
}

func DeleteBook(c *gin.Context) {
	// DELETE request body
	var input models.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	collection := database.Collection("books")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// delete book
	// check if bookid exists
	filter := bson.M{"bookid": input.BOOKID}
	cursor, err := collection.Find(ctx, filter)
	if cursor.RemainingBatchLength() == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bookid does not exist"})
		return
	}
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	filter = bson.M{"bookid": input.BOOKID}
	deletedBook := collection.FindOneAndDelete(ctx, filter)
	var book models.Book
	deletedBook.Decode(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}
