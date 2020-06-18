package controllers

import (
	"day4/models"

	"github.com/gin-gonic/gin"
)

func CreateCar(context *gin.Context) {
	type carModel struct {
		Name  string `json:"name"`
		Year  int    `json:"year"`
		Color string `json:"color"`
	}

	// Validate input
	var input carModel
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(500, gin.H{"success": false, "result": nil, "errorMessage": err.Error()})
		return
	}

	// Create todo
	car := models.Car{Name: input.Name, Year: input.Year, Color: input.Color}
	models.DB.Create(&car)

	context.JSON(201, gin.H{"success": true, "result": car, "errorMessage": nil})
}
