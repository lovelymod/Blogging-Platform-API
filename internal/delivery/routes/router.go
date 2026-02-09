package routes

import (
	"blogging-platform-api/internal/delivery/routes/handler"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	BlogHandler *handler.BlogHandler
}

func SetupRouter(r *gin.Engine, h *Handlers) {
	// Public Routes
	r.POST("/blog", h.BlogHandler.Create)

	// Protected Routes (ตัวอย่างถ้ามี Middleware)
	// api.Use(middleware.AuthMiddleware()).POST("/blogs", blogHandler.Create)
}
