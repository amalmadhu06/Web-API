//middlewares are used to secure the routes

package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"jwt_auth/initializers"
	"jwt_auth/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	fmt.Println("Check completed by middleware	")

	// 1.Get the cookie off the request

	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		fmt.Println("first error")
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	// 2.Decode / validate it

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		// 3.Check the exp

		//checking the expiry of the token
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			fmt.Println("second error")
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// 4. Find the user with token sub

		var user models.User
		initializers.DB.First(&user, claims["sub"])

		//if no user is found against the given id,
		if user.ID == 0 {
			fmt.Println("third error")
			c.AbortWithStatus(http.StatusUnauthorized)

		}

		// 5. Attach to req

		// sending the response with user object
		fmt.Println(user)
		c.Set("user", user)

		// 6.Continue
		c.Next()

	} else {
		fmt.Println("fourth error")
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
