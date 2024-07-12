package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huynhtrongtien/dove/pkg/http/request"
	"github.com/huynhtrongtien/dove/pkg/jwt"
	"github.com/huynhtrongtien/dove/services"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	SUCCESS               = 0
	UNAUTHORIZED          = -1000
	TOKEN_EXPIRED         = -2000
	AUTH_CHECK_TOKEN_FAIL = -3000
	INVALID_TOKEN         = -4000
)

const (
	KeyTokenData = "token_data"
)

// JWT is jwt middleware
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int = SUCCESS
		var payload interface{}

		bearerToken := c.GetHeader("Authorization")
		token, _ := request.ParseTokenFromBearerToken(bearerToken)

		if token == "" {
			code = UNAUTHORIZED
		} else {
			var err error
			// try to verify by JWT
			payload, err = services.JWTMaker.VerifyToken(token)
			if err != nil {
				code = AUTH_CHECK_TOKEN_FAIL
				if err == jwt.ErrInvalidToken {
					code = INVALID_TOKEN
				} else if err == jwt.ErrExpiredToken {
					code = TOKEN_EXPIRED
				}
			}
		}

		if code != SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
			})

			c.Abort()
			return
		}
		c.Set(KeyTokenData, payload)
		c.Next()
	}
}

func ParseToken(c *gin.Context) (int64, int64, error) {
	tokenData, exist := c.Get(KeyTokenData)
	if !exist {
		st, _ := status.New(codes.Unauthenticated, "Unauthenticated request").WithDetails(
			&errdetails.LocalizedMessage{
				Locale:  "vi",
				Message: "Tài khoản chưa được đăng nhập",
			},
		)

		return 0, 0, st.Err()
	}

	return tokenData.(*jwt.Payload).UserID, tokenData.(*jwt.Payload).OrganizationId, nil
}
