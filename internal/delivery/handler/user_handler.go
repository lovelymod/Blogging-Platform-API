package handler

import (
	"blogging-platform-api/internal/entity"
	"blogging-platform-api/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	usecase entity.UserUsecase
}

func NewUserHandler(usecase entity.UserUsecase) entity.UserHandler {
	return &userHandler{
		usecase: usecase,
	}
}

func (h *userHandler) Register(c *gin.Context) {
	var req entity.UserRegisterReq

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &entity.Resp{
			Message: err.Error(),
			Success: false,
		})
		return
	}

	err := h.usecase.Register(&req)

	if err != nil {
		log.Println(err)
		errHttpStatus := utils.GetHttpErrStatus(err)
		c.JSON(errHttpStatus, &entity.Resp{
			Message: err.Error(),
			Success: false,
		})
		return
	}

	c.JSON(http.StatusCreated, &entity.Resp{
		Message: "user_created",
		Success: true,
	})

}
