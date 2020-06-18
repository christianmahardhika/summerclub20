package main

import (
	"day4/controllers"
	"day4/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.ConnectDatabase()

	router.POST("/api/car/create", controllers.CreateCar)
	router.GET("/api/car/list", controllers.FindCars)
	router.PATCH("/api/car/edit", controllers.UpdateCar)
	router.DELETE("/api/car/remove", controllers.DeleteCar)

	router.Run()
}
