package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/huynhtrongtien/dove/entities"
	"github.com/huynhtrongtien/dove/pkg/jwt"
)

func IsAllow(roles ...entities.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenData, exist := c.Get("token_data")
		if !exist {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*5))
		defer cancel()

		userID := tokenData.(*jwt.Payload).UserID
		data, err := userModel.First(ctx, map[string]any{"id": userID})
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			c.Abort()
		}

		for _, val := range roles {
			if data.Role == val {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, nil)
		c.Abort()
	}
}
