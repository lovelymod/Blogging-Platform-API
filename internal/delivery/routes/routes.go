package routes

import (
	"blogging-platform-api/internal/delivery/routes/handler"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	BlogHandler *handler.BlogHandler
}

func SetupRoutes(r *gin.Engine, h *Handlers) {
	// Public Routes
	r.POST("/blog", h.BlogHandler.Create)
	r.GET("/blogs", h.BlogHandler.GetAll)

	// Protected Routes (ตัวอย่างถ้ามี Middleware)
	// api.Use(middleware.AuthMiddleware()).POST("/blogs", blogHandler.Create)
}
