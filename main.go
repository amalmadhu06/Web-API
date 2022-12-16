package main

import (
	"jwt_auth/controllers"
	"jwt_auth/initializers"
	"jwt_auth/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	//user routes
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.GET("logout", controllers.Logout)

	// admin routes

	r.POST("/admin/createAdmin", controllers.CreateAdmin)
	r.POST("/admin/login", controllers.AdminLogin)
	r.GET("/admin/logout", controllers.AdminLogout)

	r.Run()

}
