package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)


type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author 	 string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []Book{
	{ ID: "1", Title: "book 1", Author: "author 1", Quantity: 2},
	{ ID: "2", Title: "book 2", Author: "author 2", Quantity: 2},
	{ ID: "3", Title: "book 3", Author: "author 3", Quantity: 2},
}



// find book by ID [helper func]
func findBookById(id string) (*Book, error) {

	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found.")






}


// get all books
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// get book by id
func getBooksById(c *gin.Context) {
	id := c.Param("id")

	book, err := findBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": "book not found."})
		return
	}
	
	c.IndentedJSON(http.StatusOK, book)
}

// create a new book
func createBook(c *gin.Context) {
	var newBook Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}


// find a book by query params
func findBookByQuery(c * gin.Context) {

	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": "missing id query params."})
		return
	}

	books, err := findBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": "book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, books)

}


func main() {
	router := gin.Default()

	router.GET("/api/books", getBooks)
	router.GET("/api/books/:id", getBooksById)
	router.POST("/api/books", createBook)
	router.PATCH("/api/books", findBookByQuery)

	router.Run()


}
