package middleware

import (
	"fmt"
	"jwt_auth/initializers"
	"jwt_auth/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuthAdmin(c *gin.Context) {
	// 1. Get the token off the request body

	tokenString, err := c.Cookie("AdminAuth")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// 2. decode and validate jwt

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing in method : %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		// 3. checking the validity of the token
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// 4. Find the user with token sub

		var admin models.Admin
		initializers.DB.First(&admin, claims["sub"])

		// if admin is not found againist the id

		if admin.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// 5 Attaching to the response
		c.Set("admin", admin)

		// 6. Continue

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
