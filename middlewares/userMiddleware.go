package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kylerequez/go-gin-api/services"
)

func UserAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(":::-:::\tAUTHORIZATION MIDDLEWARE\t:::-:::")

		cookie, err := c.Cookie("go-gin-crud-token")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "The auth cookie does not exists",
			})
			c.Abort()
			return
		}

		isValidated, err := services.ValidateToken(cookie)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "There was an error in validating the token",
			})
			c.Abort()
			return
		}

		claims, ok := isValidated.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "There was an error in parsing the token",
			})
			return
		}

		exp := claims["expiration"].(float64)
		if int64(exp) < time.Now().Local().Unix() {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "The token has expired. Please login again.",
			})
			c.Abort()
			return
		}

		isAuthorized := claims["authorized"].(bool)
		if !isAuthorized {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "The token has expired. Please login again.",
			})
			c.Abort()
			return
		}
		fmt.Println(":::-:::\tAUTHORIZATION SUCCESSFUL\t:::-:::")
		c.Next()
	}
}
