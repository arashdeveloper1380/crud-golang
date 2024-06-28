package main

import (
	UserController "crud-api/app/controllers/user"
	"crud-api/app/models"
	UserRepository "crud-api/app/repositories/user"
	UserService "crud-api/app/services/user"
	"crud-api/packages/config"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config.Set()
	config := config.Get()

	dsn := config.DB.DB_User + ":" + config.DB.DB_Pass + "@tcp(" + config.DB.DB_Host + ":" + config.DB.DB_Port + ")/" + config.DB.DB_Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&models.User{})

	userRepository := UserRepository.Instance(db)
	userService := UserService.Instance(userRepository)
	userController := UserController.Instance(userService)

	r := gin.Default()

	r.POST("/users", userController.CreateUser)
	r.GET("/users", userController.GetAllUsers)
	r.GET("/users/:id", userController.GetUserByID)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)

	r.Run(":8080")

}
