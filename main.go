package main

import (
	"log"
	"pustaka-api/data"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/crud?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db connection error")
	}

	db.AutoMigrate(&data.Nama{})

	dataRepository := data.NewRepository(db)
	userService := data.NewService(dataRepository)
	inputHandler := handler.NewInputHandler(userService)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.POST("/name", inputHandler.PostNameHandler)
	v1.GET("/names", inputHandler.GetNames)
	v1.GET("/names/:id", inputHandler.GetName)
	v1.DELETE("/delete/:id", inputHandler.DeleteBook)
	v1.PUT("/update/:id", inputHandler.UpdateName)

	router.Run(":8080")
}
