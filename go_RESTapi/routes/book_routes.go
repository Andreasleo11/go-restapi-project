package routes

import (
	"go_RESTapi/controllers"

	"github.com/gin-gonic/gin"
)

func SetupBookRoutes(r *gin.Engine) {
	r.GET("/books", controllers.GetBooks)
	r.POST("/books", controllers.CreateBook)
	r.PUT("/books/:id", controllers.UpdateBook)    // Update Book
	r.DELETE("/books/:id", controllers.DeleteBook) // Delete Book
}
