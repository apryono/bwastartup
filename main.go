package main

import (
	"fmt"
	"log"
	"project/bwastartup/auth"
	"project/bwastartup/handler"
	"project/bwastartup/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "user=postgres password=4pry0n0 dbname=bwastartup port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()

	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.LoginUser)
	api.POST("/check_email", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)

	if err := router.Run(":8081"); err != nil {
		log.Println(err)
		return
	}
}
