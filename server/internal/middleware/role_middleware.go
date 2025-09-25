package middleware

import (
	"NWUCA-Management-System/server/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RoleAuthMiddleware 创建一个检查用户角色的中间件
func RoleAuthMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("userRole")
		if !exists {
			c.JSON(http.StatusForbidden, dto.Response{
				Code: http.StatusForbidden,
				Msg:  "用户角色不存在",
				Data: nil,
			})
			c.Abort()
			return
		}

		role, ok := userRole.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, dto.Response{
				Code: http.StatusInternalServerError,
				Msg:  "服务器内部错误",
				Data: nil,
			})
			c.Abort()
			return
		}

		if role != requiredRole {
			c.JSON(http.StatusForbidden, dto.Response{
				Code: http.StatusForbidden,
				Msg:  "权限不足",
				Data: nil,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
