package middleware

import (
	"blogging-platform-api/internal/entity"
	"blogging-platform-api/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(config *entity.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		atk := c.GetHeader("Authorization")
		splitAtk := strings.Split(atk, " ")

		if len(splitAtk) != 2 || splitAtk[0] != "Bearer" || splitAtk[1] == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, &entity.Resp{
				Message: entity.ErrAuthTokenInvalid.Error(),
				Success: false,
			})
			return
		}

		claims, err := utils.ParseAccessToken(splitAtk[1], config.ACCESS_TOKEN_SECRET)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, &entity.Resp{
				Message: err.Error(),
				Success: false,
			})
			return
		}

		c.Set("userID", claims.Subject)
		c.Next()
	}
}
