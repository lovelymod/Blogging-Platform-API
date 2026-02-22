package handler

import (
	"blogging-platform-api/internal/entity"
	"blogging-platform-api/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	usecase entity.AuthUsecase
}

func NewAuthHandler(usecase entity.AuthUsecase) entity.AuthHandler {
	return &authHandler{
		usecase: usecase,
	}
}

func (h *authHandler) Register(c *gin.Context) {
	var req entity.AuthRegisterReq

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

func (h *authHandler) Login(c *gin.Context) {
	var req entity.AuthLoginReq

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &entity.Resp{
			Message: err.Error(),
			Success: false,
		})
		return
	}

	loginResp, err := h.usecase.Login(&req)

	if err != nil {
		log.Println(err)
		errHttpStatus := utils.GetHttpErrStatus(err)
		c.JSON(errHttpStatus, &entity.Resp{
			Message: err.Error(),
			Success: false,
		})
		return
	}

	c.SetCookie("refreshToken", loginResp.RefreshToken, 60*5, "/", "localhost", false, true)
	c.JSON(http.StatusOK, &entity.Resp{
		Data: &entity.AuthLoginResp{
			AccessToken: loginResp.AccessToken,
			User:        loginResp.User,
		},
		Message: "user_logged_in",
		Success: true,
	})
}

func (h *authHandler) Logout(c *gin.Context) {
	refreshToken, err := c.Cookie("refreshToken")
	if err != nil {
		c.JSON(http.StatusUnauthorized, &entity.Resp{
			Message: err.Error(),
			Success: false,
		})
		return
	}

	if err := h.usecase.Logout(refreshToken); err != nil {
		httpErrStatus := utils.GetHttpErrStatus(err)
		c.JSON(httpErrStatus, &entity.Resp{
			Message: err.Error(),
			Success: false,
		})
		return
	}

	c.SetCookie("refreshToken", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, &entity.Resp{
		Message: "logout",
		Success: true,
	})
}

func (h *authHandler) RefreshToken(c *gin.Context) {
	rtk, err := c.Cookie("refreshToken")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, &entity.Resp{
			Message: err.Error(),
			Success: false,
		})
		return
	}

	newAtk, newRtk, err := h.usecase.RefreshToken(rtk)
	if err != nil {
		log.Println(err)
		httpErrStatus := utils.GetHttpErrStatus(err)
		c.JSON(httpErrStatus, &entity.Resp{
			Message: err.Error(),
			Success: false,
		})
		return
	}

	c.SetCookie("refreshToken", newRtk, 60*5, "/", "localhost", false, true)
	c.JSON(http.StatusOK, &entity.Resp{
		Data: &entity.AuthLoginResp{
			AccessToken: newAtk,
		},
		Success: true,
		Message: "refreshed",
	})

}
