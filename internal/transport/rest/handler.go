package rest

import (
	"context"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "my-crud-app/docs"
	"my-crud-app/internal/domain"
)

type Books interface {
	GetBook(ctx context.Context, id int) (domain.Book, error)
	GetBooks(ctx context.Context) ([]domain.Book, error)
	CreateBook(ctx context.Context, book domain.CreateBookInput) error
	UpdateBook(ctx context.Context, id int, book domain.CreateBookInput) error
	DeleteBook(ctx context.Context, id int) error
}

type Users interface {
	SignUp(ctx context.Context, inp domain.SignUpUser) error
	SignIn(ctx context.Context, inp domain.SignInUser) (string, error)
	ParseToken(ctx context.Context, token string) (int64, error)
}

type Handler struct {
	booksService Books
	userService  Users
}

func NewHandler(books Books, users Users) *Handler {
	return &Handler{
		booksService: books,
		userService:  users,
	}
}

func (h *Handler) CreateRouter() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("", h.SignUp)
		auth.GET("", h.SignIn)
	}

	books := router.Group("/books", h.userIdentity)
	{
		books.GET("/:id", h.GetBook)
		books.GET("", h.GetBooks)
		books.POST("", h.CreateBook)
		books.PUT("/:id", h.UpdateBook)
		books.DELETE("/:id", h.DeleteBook)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
