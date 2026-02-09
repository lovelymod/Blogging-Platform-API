package handler

import (
	"blogging-platform-api/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogHandler struct {
	Usecase entity.BlogUsecase
}

func (h *BlogHandler) Create(c *gin.Context) {
	var blog entity.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Usecase.Create(c.Request.Context(), &blog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, blog)
}
