package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireAuth(c *gin.Context) {
	fmt.Println("Check completed by middleware	")
	
	// 1.Get the cookie off the request

	tokenString, err := c.Cookie("Authorization")

	if err != nil{
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	
	// 2.Decode / validate it

	

	
	// 3.Check the exp
	
	// 4. Find the user with token sub
	
	// 5. Attach to req
	
	// 6.Continue
	c.Next()

}
