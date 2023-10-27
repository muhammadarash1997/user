package main

import (
	"log"

	"github.com/gin-gonic/gin"
	authcontroller "github.com/muhammadarash1997/user/auth/controller"
	authservice "github.com/muhammadarash1997/user/auth/service"
	"github.com/muhammadarash1997/user/database"
	usercontroller "github.com/muhammadarash1997/user/user/controller"
	userrepository "github.com/muhammadarash1997/user/user/repository"
	userservice "github.com/muhammadarash1997/user/user/service"
)

func main() {
	r := gin.Default()
	r.Use(authservice.CORSMiddleware())

	db, err := database.StartConnection()
	if err != nil {
		log.Printf("Error :%v", err)
		return
	}

	userRepository := userrepository.NewRepository(db)

	userService := userservice.NewService(userRepository)
	authService := authservice.NewService(userRepository)

	userController := usercontroller.NewController(userService)
	authController := authcontroller.NewController(authService, userService)

	r.POST("/login", authController.LoginHandler)
	r.POST("/register", userController.RegisterUserHandler)
	r.GET("/users", authController.AuthenticateHandler, userController.GetUserHandler)

	r.Run(":8080")
}
