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

	r.POST("/admin/login", controllers.AdminLogin)
	r.GET("/admin/validate", middleware.RequireAuthAdmin, controllers.AdminValidate)
	r.GET("/admin/logout", controllers.AdminLogout)

	// to be deleted
	r.POST("/admin/createAdmin", middleware.RequireAuthAdmin, controllers.CreateAdmin)

	//crud operation by admin on user

	r.GET("admin/viewAllUsers", middleware.RequireAuthAdmin, controllers.ViewAllUsers)
	r.GET("admin/viewUser", middleware.RequireAuthAdmin, controllers.ViewUser)
	r.DELETE("admin/deleteUser", middleware.RequireAuthAdmin, controllers.DeleteUser)
	r.POST("admin/createUser", middleware.RequireAuthAdmin, controllers.CreateUser)
	r.POST("admin/updateUserPassword", middleware.RequireAuthAdmin, controllers.UpdateUserPassword)

	r.Run()

}
