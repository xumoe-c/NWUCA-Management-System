package middleware

import (
	"NWUCA-Management-System/server/internal/dto"
	"NWUCA-Management-System/server/internal/util/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, dto.Response{
				Code: http.StatusUnauthorized,
				Msg:  "没有Token",
				Data: nil,
			})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, dto.Response{
				Code: http.StatusUnauthorized,
				Msg:  "Token格式不正确",
				Data: nil,
			})
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := auth.ParseToken(tokenString, secretKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.Response{
				Code: http.StatusUnauthorized,
				Msg:  "Token不正确",
				Data: nil,
			})
			c.Abort()
			return
		}

		// 将用户信息存储在下文中
		c.Set("userID", claims.UserID)
		c.Set("userRole", claims.Role)

		c.Next()
	}
}
