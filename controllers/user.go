package controllers

import (
	"jwt_auth/initializers"
	"jwt_auth/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// -------------------------------------------------FUNCTION Signup-----------------------------------------------------------------------

func Signup(c *gin.Context) {
	//1. get the email and password from the req body

	//struct for holding the values that come from request
	var body struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	//checking if body is present, if not, will respond with below status
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input, please check",
		})
		return
	}

	// 2.Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash the password",
		})
		return
	}

	// 3. Create the user

	//for this, we need to save the credentials to the databse.
	//following lines will help to do this
 
	user := models.User{Email: body.Email, Password: string(hash)}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create user",
		})
		return
	}

	// 4. respond
	// upon successful user creation, sending this back
	c.JSON(http.StatusOK, gin.H{
		"message": "user created successfully",
	})
}

// -------------------------------------------------FUNCTION Login------------------------------------------------------------------------

func Login(c *gin.Context) {
	// 1. get the email id and password from req body

	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read the body",
		})
		return
	}

	// 2. Look up requested user

	var user models.User //for storing value retrieved from the database matching with the entered email id

	// `First` finds the first record ordered by primary key, matching given conditions and store it to the given condition (first argument)
	initializers.DB.First(&user, "email = ?", body.Email)

	//if no user is found againt the given email id, then
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid username or password",
		})
		return
	}

	// 3. Compare sent in password with hashed password in the db

	//CompareHashAndPassword compares a bcrypt hashed password with its possible plaintext equivalent.
	//Returns nil on success, or an error on failure.
	//user.Password is the hashed password in the DB
	//body.Password is the password received through the request

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	// if password doesn't match
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	// 4. Create a JWT token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,                                    //subject : user.ID
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(), //setting the expiration time to 30 days
	})

	//sign and get the complete encoded token as a string using the secret key

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	// 5. send it back

	//sending the token as cookies
	c.SetSameSite(http.SameSiteLaxMode)

	// SetCookie arguments => name of the cookie : string, value of the cookie : string,
	// Expiration time : int, path : string, domain name : string, secure : bool, httpOnly : bool
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		// directly sending the token as response (not recommneded, so commenting it)
		// "token": tokenString,
		"message": "Logged in successfully",
	})
}

// -------------------------------------------------FUNCTION Logout------------------------------------------------------------------------
func Logout(c *gin.Context) {
	// if cookie is present, set the expiry to -1 and direct to login page
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	// c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "logged out successfully",
	})
}

// ----------------------------------------------------------FUNCTION validate -----------------------------------------------------------------------

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

// ----------------------------------------------------------FUNCTION homePage -----------------------------------------------------------------------

func HomePage(c *gin.Context) {

	// 1. check if the user is already logged in
	//    requireAuth.go
	// 2. if yes, give access to homepage

	// 3. else, route to login page

}
