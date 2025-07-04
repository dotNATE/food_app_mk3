package middleware

import (
	"main/pkg/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	jwt_secret := []byte(os.Getenv("JWT_SECRET"))

	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("token")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.CreateErrorHTTPResponse("Missing or invalid token header", nil))
			return
		}

		token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, jwt.ErrTokenUnverifiable
			}
			return jwt_secret, nil
		})

		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.CreateErrorHTTPResponse("Invalid or expired token", nil))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			user_id := claims["user_id"]

			ctx.Set("user_id", user_id)
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.CreateErrorHTTPResponse("Invalid token claims", nil))
		}
	}
}
