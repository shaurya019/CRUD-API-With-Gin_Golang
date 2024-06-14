package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type Book struct {
	ID     string `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

var books = []Book{
	{ID: "1", Title: "1984", Author: "George Orwell"},
    {ID: "2", Title: "Brave New World", Author: "Aldous Huxley"},
}

func getBooks(c *gin.Context){
	c.JSON(http.StatusOK,books)
}

func getBookByID(c *gin.Context) {
    id := c.Param("id")
	// id := c.Param("id"): Retrieves the id parameter from the URL.
    for _, book := range books {
        if book.ID == id {
            c.JSON(http.StatusOK, book)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func addBook(c *gin.Context){
	var x Book
	if err := c.ShouldBindBodyWithJSON(&x); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}

	books = append(books, x)
	c.JSON(http.StatusCreated,x)
}

// var newBook Book: Declares a new Book variable.
// c.ShouldBindJSON(&newBook): Binds the JSON body of the request to newBook. If there's an error, responds with HTTP status code 400 (Bad Request) and the error message.


func main() {
	r:=gin.Default();

	r.GET("/books", getBooks)
    r.GET("/books/:id", getBookByID)
    r.POST("/books", addBook)

    r.Run(":8080")
}