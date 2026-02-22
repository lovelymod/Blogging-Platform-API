package router

import (
	"blogging-platform-api/internal/entity"
	"blogging-platform-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	BlogHandler entity.BlogHandler
	AuthHandler entity.AuthHandler
}

func SetupRoutes(r *gin.Engine, h *Handlers, config *entity.Config) {
	api := r.Group("/api")

	// Auth
	{
		api.POST("/register", h.AuthHandler.Register)
		api.POST("/login", h.AuthHandler.Login)
		api.POST("/refresh-token", h.AuthHandler.RefreshToken)
	}

	protected := api.Use(middleware.AuthMiddleware(config))

	// Auth
	{
		protected.POST("/logout", h.AuthHandler.Logout)
	}

	// Blog
	{
		protected.GET("/blogs", h.BlogHandler.GetAll)
		protected.GET("/blog/:id", h.BlogHandler.GetByID)
		protected.POST("/blog", h.BlogHandler.Create)
		protected.PUT("/blog/:id", h.BlogHandler.Update)
		protected.DELETE("/blog/:id", h.BlogHandler.Delete)
	}

}
