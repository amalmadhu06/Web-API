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

// ----------------------FUNCTION CreateAdmin---------------

func CreateAdmin(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read the request body",
		})
		return
	}

	//hashing the password

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash the password",
		})
		return
	}

	//creating admin

	// saving credentials to the database

	admin := models.Admin{Email: body.Email, Password: string(hash)}

	result := initializers.DB.Create(&admin)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create admin",
		})
		return
	}

	//sending response

	c.JSON(http.StatusOK, gin.H{
		"message": "admin created successfully",
	})

}

// --------------------------------------FUNCTION adminLogin------------------------------------------------------------
func AdminLogin(c *gin.Context) {
	// 1. get email id and password from the request header
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
	// 2. Look up for the requested admin
	var admin models.Admin

	initializers.DB.First(&admin, "email = ?", body.Email)

	if admin.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid username of password",
		})
		return
	}
	// 3. Compare sent in password with hashed password in the db

	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	// 4. Create a JWT token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": admin.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	// 5. send it back

	c.SetSameSite(http.SameSiteLaxMode)

	c.SetCookie("AdminAuth", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

// --------------------------------------FUNCTION AdminLogout------------------------------------------------------------
// func AdminLogout() {

// }

// --------------------------------------FUNCTION deleteUser------------------------------------------------------------
// func DeleteUser() {

// }

// --------------------------------------FUNCTION addUser------------------------------------------------------------
// func AddUser() {

// }

// --------------------------------------FUNCTION adminLogin------------------------------------------------------------
// func UpdateUser() {

// }