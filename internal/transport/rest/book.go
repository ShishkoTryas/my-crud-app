package rest

import (
	"context"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"my-crud-app/internal/domain"
	"net/http"
	"strconv"
)

// @Summary GetBook
// @Description  Get one book
// @Tags         Books
// @Accept       json
// @Produce      json
// @Param        id path string true "Book ID"
// @Success      200
// @Failure      404
// @Router       /books/{id} [get]
func (h *Handler) GetBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Error during conversion")
		return
	}
	book, err := h.booksService.GetBook(context.TODO(), id)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "GetBook",
		}).Error(err)
		return
	}

	c.JSON(http.StatusOK, book)
}

// @Summary GetBooks
// @Description  Get all books
// @Tags         Books
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      404
// @Router       /books [get]
func (h *Handler) GetBooks(c *gin.Context) {
	books, err := h.booksService.GetBooks(context.TODO())
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "GetBooks",
		}).Error(err)
		return
	}
	c.JSON(http.StatusOK, books)
}

// @Summary CreateBook
// @Description  Create new book
// @Tags         Books
// @Accept       json
// @Produce      json
// @Param        input body domain.CreateBookInput true "create"
// @Success      200
// @Failure      404
// @Router       /books [post]
func (h *Handler) CreateBook(c *gin.Context) {
	var inputData domain.CreateBookInput
	if err := c.BindJSON(&inputData); err != nil {
		return
	}
	err := h.booksService.CreateBook(context.TODO(), inputData)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "createBook",
			"problem": "reading request body",
		}).Error(err)
		return
	}
	c.JSON(http.StatusOK, inputData)
}

// @Summary UpdateBook
// @Description  Update book by id
// @Tags         Books
// @Accept       json
// @Produce      json
// @Param        id path string true "Book ID"
// @Param        input body domain.CreateBookInput true "create"
// @Success      200
// @Failure      404
// @Router       /books/{id} [put]
func (h *Handler) UpdateBook(c *gin.Context) {
	var inputData domain.CreateBookInput
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Error during conversion")
		return
	}
	c.BindJSON(&inputData)
	err = h.booksService.UpdateBook(context.TODO(), id, inputData)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "updateBook",
			"problem": "reading request body",
		}).Error(err)
		return
	}
	c.JSON(http.StatusOK, inputData)
}

// @Summary DeleteBook
// @Description  Delete book by id
// @Tags         Books
// @Accept       json
// @Produce      json
// @Param        id path string true "Book ID"
// @Success      200
// @Failure      404
// @Router       /books/{id} [delete]
func (h *Handler) DeleteBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Error during conversion")
		return
	}
	err = h.booksService.DeleteBook(context.TODO(), id)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "deleteBook",
		}).Error(err)
		return
	}
}
